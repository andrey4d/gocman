package main

import (
	"encoding/json"
	"fmt"
	"godman/internal/containers"
	"godman/internal/helpers"
	"log"
	"os"
	"strings"

	"github.com/google/go-containerregistry/pkg/crane"
	"golang.org/x/exp/maps"

	// v1 "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

const imageStore = "containers/images"

type manifest []struct {
	Config   string
	RepoTags []string
	Layers   []string
}

type ImageRow map[string]string  // imageName: tag
type ImageDB map[string]ImageRow // {id : {"imageName":"tag"}}

// DB := ImageDB{
// 	"a416a98b71e224a3": ImageRow{
// 		"busybox": "latest",
// 	},
// }

func main() {
	fmt.Println("Image parse")

	// image := containers.GetAbsPath("busybox.tar")

	// getImageManifest("busybox")
	// mkImage()

	// saveImageDB(DB)

	addImageToDB("busybox", "latest", "a416a98b71e224a3")
	addImageToDB("busybox", "3.18", "a416a98b71e23ed8")
	addImageToDB("registry.home.local/busybox", "latest", "a416a98b71e224a3")
	listImages()
	// name, tag := getImageNameAndTagByShaHex("a416a98b71e224a3")
	// fmt.Println(name, ":", tag)
	// id, err := removeImageFromDbByHash("a416a98b71e23ed8")
	// id, err := getIdByName("busybox:3.18")
	// id, err := getIdByName("registry.home.local/busybox")
	// id, err := removeImageFromDbByName("busybox:3.18")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(id)
	id, err := removeImageFromDbByName("registry.home.local/busybox")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
	id, err = removeImageFromDbByName("busybox")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

}

func listImages() {
	imagesDB := loadImageDB()
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("NAME", "TA", "ID")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for shaHash, meta := range imagesDB {
		for image, tag := range meta {
			tbl.AddRow(image, tag, shaHash)
		}
	}
	tbl.Print()
}

func getImageNameAndTagById(shaHex string) (string, string) {
	imageDB := loadImageDB()
	if imageDB[shaHex] != nil {
		row := imageDB[shaHex]
		key := maps.Keys(row)[0] // get first image
		return key, row[key]
	}
	return "", ""
}

func getIdByName(name string) (string, error) {
	imagesDB := loadImageDB()

	name, tag := getImageNameAndTag(name)
	for id, meta := range imagesDB {
		for n, t := range meta {
			if n == name && t == tag {
				return id, nil
			}
		}
	}
	return "", fmt.Errorf("no ID associated with name %s", name)
}

func removeImageFromDbByHash(id string) (string, error) {
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

func removeImageFromDbByName(name string) (string, error) {
	imagesDB := loadImageDB()
	id, err := getIdByName(name)
	if err != nil {
		return "", fmt.Errorf("image %s not found", name)
	}
	name, tag := getImageNameAndTag(name)
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

func saveImageDB(imageDB ImageDB) {
	imageDbPath := containers.GetAbsPath("containers") + "/images/images.json"

	imagesBytes, err := json.Marshal(imageDB)
	helpers.ErrorHelperPanicWithMessage(err, "marshal imageDB to JSON")
	helpers.ErrorHelperPanicWithMessage(os.WriteFile(imageDbPath, imagesBytes, 0644), "write imageDB to images.json")
}

func loadImageDB() ImageDB {
	imageDbPath := containers.GetAbsPath("containers") + "/images/images.json"

	if _, err := os.Stat(imageDbPath); os.IsNotExist(err) {
		helpers.ErrorHelperPanicWithMessage(os.WriteFile(imageDbPath, []byte(`{}`), 0644), "make images.json")
	}

	file, err := os.Open(imageDbPath)
	if err != nil {
		helpers.FatalHelperLog("read images.json")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var imageDB ImageDB
	helpers.ErrorHelperPanicWithMessage(decoder.Decode(&imageDB), "unable decode images.json")

	return imageDB
}

func getImageManifest(name string) {

	imgName, tagName := getImageNameAndTag(name)

	log.Printf("Downloading metadata for %s:%s, please wait...", imgName, tagName)
	img, err := crane.Pull(strings.Join([]string{imgName, tagName}, ":"))
	if err != nil {
		log.Fatal(err)
	}

	manifest, _ := img.Manifest()
	imageShaHex := manifest.Config.Digest.Hex[:16]
	tmp := containers.GetTempPath(containers.GetAbsPath("tmp"))
	log.Println(tmp)
	log.Printf("imageHash: %v\n", imageShaHex)
	log.Println("Checking if image exists under another name...")
	// log.Printf("%v", manifest)
	helpers.ErrorHelperPanicWithMessage(crane.Save(img, imageShaHex, fmt.Sprintf("%s/%s.tar", tmp, name)), "save image")

}

// func downloadImageMetadata()

// func downloadImage(img v1.Image, imageShaHex string, src string) {
// 	path := containers.GetAbsPath("tmp") + "/" + imageShaHex
// 	os.Mkdir(path, 0755)
// 	path += "/package.tar"
// 	/* Save the image as a tar file */
// 	if err := crane.SaveLegacy(img, src, path); err != nil {
// 		log.Fatalf("saving tarball %s: %v", path, err)
// 	}
// 	log.Printf("Successfully downloaded %s\n", src)
// }

func getImageNameAndTag(src string) (string, string) {
	s := strings.Split(src, ":")
	if len(s) > 1 {
		return s[0], s[1]
	} else {
		return s[0], "latest"
	}
}

func mkImage() {
	c := map[string][]byte{
		"/binary": []byte("binary contents"),
	}
	i, _ := crane.Image(c)
	d, _ := i.Digest()
	fmt.Println(d)
	m, _ := i.Manifest()
	fmt.Println(m)
}
