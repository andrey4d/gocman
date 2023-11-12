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

type ImageRow map[string]string  // tag: shaHash
type ImageDB map[string]ImageRow // {imageName: {"tag":"shaHash"}}
// DB := ImageDB{
// 	"busybox": ImageRow{
// 		"latest": "a416a98b71e224a3",
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
}

func listImages() {
	imagesDB, _ := loadImageDB()
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("NAME", "TA", "ID")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for image, meta := range imagesDB {
		for tag, shaHash := range meta {
			tbl.AddRow(image, tag, shaHash)
		}
	}
	tbl.Print()
}

func removeImageFromDB() {
	imagesDB, _ := loadImageDB()
	fmt.Println(imagesDB)
}

func addImageToDB(imageName string, tag string, shaHash16 string) {
	imageDB, _ := loadImageDB()
	row := ImageRow{}
	if imageDB[imageName] != nil {
		row = imageDB[imageName]
	}

	row[tag] = shaHash16
	imageDB[imageName] = row
	saveImageDB(imageDB)
}

func saveImageDB(imageDB ImageDB) {
	imageDbPath := containers.GetAbsPath("containers") + "/images/images.json"

	file, err := os.OpenFile(imageDbPath, os.O_WRONLY, 0644)
	helpers.ErrorHelperPanicWithMessage(err, "write imageDB to images.json")
	encoder := json.NewEncoder(file)

	helpers.ErrorHelperPanicWithMessage(encoder.Encode(imageDB), "encode imageDB")
	defer file.Close()

}

func loadImageDB() (ImageDB, error) {
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

	return imageDB, nil

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
