/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package main

import (
	"fmt"
	"godman/cmd/test/config"
	"path/filepath"
)

var Context config.Cfg

type Pointer struct {
	x int32
	y int32
}

type Position struct {
	*Pointer
}

func (p *Pointer) Move(x, y int32) {
	p.x += x
	p.y += y
}

func (p *Pointer) MoveFast(x, y int32) {
	p.x *= x
	p.y *= y
}

func NewPosition() *Position {
	return &Position{
		Pointer: &Pointer{},
	}
}

func init() {
	Context = *config.NewCfg()
	Context.SetPath("any_folder")
	Context.SetTemp("/tmp")
}

func main() {

	position := NewPosition()
	position.Pointer.Move(1, 3)
	fmt.Println(position.Pointer)

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
