/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package main

import (
	"fmt"

	// "godman/internal/config"

	"github.com/andrey4d/gocman/internal/containers"

	"github.com/google/go-containerregistry/pkg/crane"
)

const imageStore = "containers/images"
const temp = "tmp"

func main() {
	fmt.Println("Image parse")

	// image := containers.GetAbsPath("busybox.tar")

	// getImageManifest("busybox")
	// mkImage()

	// saveImageDB(DB)
	// cfg := config.InitConfig("config/config.yaml")
	// addImageToDB("busybox", "latest", "a416a98b71e224a3")
	// addImageToDB("busybox", "3.18", "a416a98b71e23ed8")
	// addImageToDB("registry.home.local/busybox", "latest", "a416a98b71e224a3")
	// containers.ListImages()
	// id, err := removeImageFromDbByHash("a416a98b71e23ed8")
	// id, err := getIdByName("busybox:3.18")
	// id, err := getIdByName("registry.home.local/busybox")
	// id, err := removeImageFromDbByName("registry.home.local/busybox")
	// if err != nil {
	// fmt.Println(err)
	// }
	// fmt.Println(id)
	// containers.DownloadImage("registry.home.local/busybox")
	// containers.DownloadImage("gcr.io/kubernetes-e2e-test-images/echoserver:2.2", cfg.Container.TempPath)
	// getManifest("alpine")

	getLowerLayers("4081d9a831083d9e57c49a95632feaf0103bd4db2c9fa1e01b48b7b1136a946d")
}

func getLowerLayers(id string) []string {
	destConfig := fmt.Sprintf("%s/%s/manifest.json", "/home/andrey/git_project/golang/godman/containers/storage/overlay-images", id)

	manifest := containers.GetManifest(destConfig)
	for _, layer := range manifest[0].Layers {
		fmt.Println(layer)
	}
	return []string{""}
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

// func getManifest(name string) {
// 	imgName, tagName := getImageNameAndTag(name)
// 	log.Printf("Downloading metadata for %s:%s, please wait...", imgName, tagName)
// 	mbytes, err := crane.Manifest(strings.Join([]string{imgName, tagName}, ":"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf(string(mbytes))
// 	var manifest *v1.Manifest
// 	log.Println(manifest)
// 	json.Unmarshal(mbytes, &manifest)
// 	id := manifest.Config.Digest.Hex
// 	log.Println(id)
// }
