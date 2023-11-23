/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package containers

import (
	"encoding/json"
	"fmt"
	"godman/internal/config"
	"godman/internal/helpers"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/fatih/color"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/rodaine/table"
	"golang.org/x/exp/maps"
)

type ImageRow map[string]string  // imageName: tag
type ImageDB map[string]ImageRow // {id : {"imageName":"tag"}}
// DB := ImageDB{
// 	"a416a98b71e224a3": ImageRow{
// 		"busybox": "latest",
// 	},
// }

type Manifest []struct {
	Config   string   `json:"Config"`
	RepoTags []string `json:"RepoTags"`
	Layers   []string `json:"Layers"`
}

func DownloadImage(name string) string {

	imgName, tagName := GetImageNameAndTag(name)
	id, _ := GetIdByName(name)
	if id == "" {
		log.Printf("Downloading metadata for %s:%s, please wait...", imgName, tagName)
		img, err := crane.Pull(strings.Join([]string{imgName, tagName}, ":"))
		helpers.CheckError(err, "pull image")

		manifest, _ := img.Manifest()
		id := manifest.Config.Digest.Hex

		log.Printf("Id: %v\n", id)
		log.Println("Checking if image exists under another name...")
		name, tag := GetImageNameAndTagById(id)
		if name != "" {
			log.Printf("The image you requested %s:%s is the same as %s:%s\n", imgName, tagName, name, tag)
		} else {
			tmp := helpers.MakeTempPath(config.Config.GetContainersTemp(), id)
			tarFile := fmt.Sprintf("%s/%s.tar", tmp, id)
			untarPath := fmt.Sprintf("%s/image.tar", tmp)

			log.Println(tarFile)
			helpers.CheckError(crane.Save(img, id, tarFile), "can't save image")
			helpers.CheckError(helpers.Untar(tarFile, untarPath), "can't untar image")

			unpackImage(id)
		}

		addImageToDB(imgName, tagName, id)
		return id

	} else {
		log.Println("Image already exists. Not downloading.")
		return id
	}
}

func GetManifest(manifestPath string) Manifest {
	// manifestPath := fmt.Sprintf("%s/%s/manifest.json", config.Config.GetOverlayImage(), id)

	manifestBytes, err := os.ReadFile(manifestPath)
	helpers.CheckError(err, "unpackImage() read manifest file")
	var manifest Manifest
	helpers.CheckError(json.Unmarshal(manifestBytes, &manifest), "unpackImage() unmarshal manifest")
	return manifest
}

func unpackImage(id string) {
	path := helpers.MakeTempPath(config.Config.GetContainersTemp(), id)

	manifest := GetManifest(fmt.Sprintf("%s/image.tar/manifest.json", path))
	configPath := fmt.Sprintf("%s/image.tar/%s", path, manifest[0].Config)
	imagesDir := fmt.Sprintf("%s/storage/overlay", config.Config.GetContainersPath())

	lowers := ""
	for _, layer := range manifest[0].Layers {
		layerDir := fmt.Sprintf("%s/%s/diff", imagesDir, layer[:64])

		if _, err := os.Stat(layerDir); os.IsNotExist(err) {

			log.Printf("Uncompressing layer to: %s \n", layerDir)
			helpers.CheckError(makeOverlayDir(imagesDir, layer[:64]), "unpackImage() can't create layer dir")

			log.Printf("Save Layer %s. \n", layer[:64])
			srcLayer := fmt.Sprintf("%s/image.tar/%s", path, layer)
			helpers.CheckError(helpers.Untar(srcLayer, layerDir), fmt.Sprintf("unpackImage() Unable to untar layer file: %s\n", srcLayer))

			helpers.CheckError(createSymlink(layerDir, &lowers), "unpackImage() can't create symlink")
		} else {
			log.Printf("Skip exist Layer %s\n", layer[:64])
		}
	}

	destConfig := fmt.Sprintf("%s/%s", config.Config.GetOverlayImage(), id)
	helpers.CheckError(helpers.MakeDirAllIfNotExists(destConfig, config.Config.GetPermissions()), "can't create  overlay-images dir")

	helpers.CheckError(helpers.Copy(configPath, fmt.Sprintf("%s/%s", destConfig, manifest[0].Config)), "unpackImage() copy config")
	helpers.CheckError(helpers.Copy(fmt.Sprintf("%s/image.tar/manifest.json", path), fmt.Sprintf("%s/manifest.json", destConfig)), "unpackImage() copy config")
}

func createSymlink(layerDir string, lower *string) error {

	linkId := helpers.GenerateID(26)
	symlinkPath := fmt.Sprintf("%s/%s", config.Config.GetOverlayLinkDir(), linkId)
	if err := os.Symlink(layerDir, symlinkPath); err != nil {
		return err
	}
	if len(*lower) > 0 {
		helpers.CheckError(os.WriteFile(layerDir[:len(layerDir)-5]+"/lower", []byte(*lower), 0644), "")
		log.Println(*lower)
	}

	if len(*lower) == 0 {
		*lower = "l/" + linkId
	} else {
		*lower += ":l/" + linkId
	}

	if err := os.WriteFile(fmt.Sprintf("%s/link", layerDir[:len(layerDir)-5]), []byte(linkId), 0644); err != nil {
		return err
	}
	return nil
}

func makeOverlayDir(imagesDir, layer string) error {
	overlays := []string{"merged", "work", "diff"}
	for _, overlay := range overlays {
		err := helpers.MakeDirAllIfNotExists(fmt.Sprintf("%s/%s/%s", imagesDir, layer, overlay), config.Config.GetPermissions())
		if err != nil {
			return err
		}
	}
	return nil
}

func GetImageNameAndTag(src string) (string, string) {
	s := strings.Split(src, ":")
	if len(s) > 1 {
		return s[0], s[1]
	} else {
		return s[0], "latest"
	}
}

func GetImageNameAndTagById(shaHex string) (string, string) {
	imageDB := loadImageDB()
	if imageDB[shaHex] != nil {
		row := imageDB[shaHex]
		key := maps.Keys(row)[0] // get first image
		return key, row[key]
	}
	return "", ""
}

func GetIdByName(name string) (string, error) {
	imagesDB := loadImageDB()

	name, tag := GetImageNameAndTag(name)
	for id, meta := range imagesDB {
		for n, t := range meta {
			if n == name && t == tag {
				return id, nil
			}
		}
	}
	return "", fmt.Errorf("no ID associated with name %s", name)
}

func GetLowerLayers(id string) []string {
	destConfig := fmt.Sprintf("%s/%s/manifest.json", config.Config.GetOverlayImage(), id)
	manifest := GetManifest(destConfig)
	layers := []string{}
	for _, layer := range manifest[0].Layers {
		layers = append(layers, strings.Split(layer, ".")[0])
	}
	return layers
}

func ListImages() {
	imagesDB := loadImageDB()
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("NAME", "TA", "IMAGE ID")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for shaHash, meta := range imagesDB {
		for image, tag := range meta {
			tbl.AddRow(image, tag, shaHash)
		}
	}
	tbl.Print()
}

func addImageToDB(imageName string, tag string, idShaHash16 string) {
	imagesDB := loadImageDB()
	row := ImageRow{}
	if imagesDB[idShaHash16] != nil {
		row = imagesDB[idShaHash16]
	}
	row[imageName] = tag
	imagesDB[idShaHash16] = row
	saveImageDB(imagesDB)
}
func RemoveImageFromDbByHash(id string) (string, error) {
	imagesDB := loadImageDB()
	row := imagesDB[id]
	if row == nil {
		return "", fmt.Errorf("no image associated with ID %s", id)
	}
	if len(row) > 1 {
		return "", fmt.Errorf("more one TAG associated with ID %s", id)
	}
	delete(imagesDB, id)
	saveImageDB(imagesDB)
	return id, nil
}

func RemoveImageFromDbByName(name string) (string, error) {
	imagesDB := loadImageDB()
	id, err := GetIdByName(name)
	if err != nil {
		return "", fmt.Errorf("image %s not found", name)
	}
	name, tag := GetImageNameAndTag(name)
	row := imagesDB[id]
	for n, t := range row {
		if n == name && t == tag {
			delete(row, name)
		}
		if len(row) == 0 {
			fmt.Println(row, len(row))
			delete(imagesDB, id)
		}
		saveImageDB(imagesDB)
	}
	return id, nil
}

func saveImageDB(imageDB ImageDB) {
	imagesBytes, err := json.Marshal(imageDB)
	helpers.CheckError(err, "marshal imageDB to JSON")
	helpers.CheckError(os.WriteFile(config.Config.GetImageDbPath(), imagesBytes, 0644), "write imageDB to images.json")
}

func loadImageDB() ImageDB {
	if _, err := os.Stat(config.Config.GetImageDbPath()); os.IsNotExist(err) {
		helpers.CheckError(os.WriteFile(config.Config.GetImageDbPath(), []byte(`{}`), 0644), "make images.json")
	}

	file, err := os.Open(config.Config.GetImageDbPath())
	if err != nil {
		helpers.FatalHelperLog("read images.json")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var imageDB ImageDB
	helpers.CheckError(decoder.Decode(&imageDB), "unable decode images.json")

	return imageDB
}
