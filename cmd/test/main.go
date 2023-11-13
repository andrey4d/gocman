package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	file_tar := "filename.tar"

	file_tgz := "filename.tar.gz"

	fmt.Println(fileTypeBuExt(file_tar))
	fmt.Println(fileTypeBuExt(file_tgz))

}

func fileTypeBuExt(filename string) string {
	ext := filepath.Ext(filename)
	t := ""
	switch ext {
	case ".gz":
		return "gzip"
	case ".tar":
		return "tar"
	case ".tgz":
		return "gz"
	}
	return t
}
