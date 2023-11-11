package main

import (
	"fmt"
	"godman/internal/containers"
	"log"
	"strings"

	"github.com/google/go-containerregistry/pkg/crane"
	// v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

const imageStore = "containers/images"

func main() {
	fmt.Println("Image parse")

	// image := containers.GetAbsPath("busybox.tar")
	tmp := containers.GetTempPath(containers.GetAbsPath("tmp"))
	log.Println(tmp)

	getImageManifest("busybox")

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
	log.Printf("imageHash: %v\n", imageShaHex)
	log.Println("Checking if image exists under another name...")
	log.Printf("%v", manifest)
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
