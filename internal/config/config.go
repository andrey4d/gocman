/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package config

import (
	"io/fs"
)

type ContainerConfig struct {
	containersPath string
	containersTemp string
	imageDbPath    string
	overlayLinkDir string
	overlayImage   string
	overlayDir     string
	permissions    fs.FileMode
}

var (
	Config ContainerConfig
)

func (c ContainerConfig) GetContainersPath() string {
	return c.containersPath
}

func (c ContainerConfig) GetContainersTemp() string {
	return c.containersTemp
}

func (c ContainerConfig) GetImageDbPath() string {
	return c.imageDbPath
}

func (c ContainerConfig) GetOverlayLinkDir() string {
	return c.overlayLinkDir
}

func (c ContainerConfig) GetOverlayImage() string {
	return c.overlayImage
}

func (c ContainerConfig) GetOverlayDir() string {
	return c.overlayDir
}

func (c ContainerConfig) GetPermissions() fs.FileMode {
	return c.permissions
}

func (c *ContainerConfig) SetContainersPath(path string) {
	c.containersPath = path
}

func (c *ContainerConfig) SetContainersTemp(path string) {
	c.containersTemp = path
}

func (c *ContainerConfig) SetImageDbPath(path string) {
	c.imageDbPath = path
}

func (c *ContainerConfig) SetOverlayLinkDir(path string) {
	c.overlayLinkDir = path
}

func (c *ContainerConfig) SetOverlayImage(path string) {
	c.overlayImage = path
}

func (c *ContainerConfig) SetOverlayDir(path string) {
	c.overlayDir = path
}

func (c *ContainerConfig) SetPermissions(perm fs.FileMode) {
	c.permissions = perm
}
