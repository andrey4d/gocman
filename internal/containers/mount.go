/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package containers

import (
	"fmt"
	"godman/internal/config"
	"godman/internal/helpers"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

type OvfsMountCfg struct {
	Lowerdir    []string // overlay/l/<ZLR4NWYDXWB5LCOBDH7WAGVYDI> --> storage/overlay/<id>/diff
	Upperdir    string   // storage/overlay/<id>/diff
	Workdir     string   // storage/overlay/<id>/work
	Target      string   // storage/overlay/<id>/merged
	Permeations fs.FileMode
	SELebel     string
}

func (o *OvfsMountCfg) getOverlayOpt(imageId string) string {
	o.Lowerdir = []string{}
	laers := GetLowerLayers(imageId)
	for _, layer := range laers {
		o.Lowerdir = append(o.Lowerdir, fmt.Sprintf("%s/%s/diff", config.Config.GetOverlayDir(), layer))
	}
	opts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s",
		strings.Join(o.Lowerdir, ":"),
		helpers.GetAbsPath(o.Upperdir),
		helpers.GetAbsPath(o.Workdir))
	return opts
}

func (o *OvfsMountCfg) createDirectoryStructure() {
	mounts := []string{
		o.Workdir,
		o.Upperdir,
		o.Target,
	}

	for _, point := range mounts {
		helpers.CheckError(helpers.MakeDirAllIfNotExists(point, o.Permeations), fmt.Sprintf("container() can't make overlay layer %s\n", point))
	}
}

func MountProc(new_root string) {
	target := filepath.Join(new_root, "/proc")
	helpers.CheckError(helpers.MakeDirAllIfNotExists(target, 0555), "MountProc() can't create /proc")
	helpers.CheckError(syscall.Mount("proc", target, "proc", 0, ""), "MountProc() can't mount /proc")
}

func UmountProc() error {
	if err := syscall.Unmount("/proc", 0); err != nil {
		return err
	}
	return nil
}

func MountRoot(new_root string) error {

	old_root := filepath.Join(new_root, "/.pivot_root")
	if err := syscall.Mount(new_root, new_root, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		helpers.ErrorHelperLog("Bind newroot")
		return err
	}

	if err := os.MkdirAll(old_root, 0700); err != nil {
		helpers.ErrorHelperLog("make dirs " + old_root)
		return err
	}

	if err := syscall.PivotRoot(new_root, old_root); err != nil {
		helpers.ErrorHelperLog("syscall.PivotRoot")
		return err
	}

	if err := syscall.Chdir("/"); err != nil {
		helpers.ErrorHelperLog("change dir to /")
		return err
	}

	if err := syscall.Unmount("/.pivot_root", syscall.MNT_DETACH); err != nil {
		helpers.ErrorHelperLog("umount " + old_root)
		return err
	}

	if err := os.RemoveAll("/.pivot_root"); err != nil {
		helpers.ErrorHelperLog("remove " + old_root)
		return err
	}

	return nil
}

func MountOvfs(imageId string, ovfs *OvfsMountCfg) error {
	fmt.Print("mount overlay fs root ...\n")
	err := syscall.Mount("overlay", ovfs.Target, "overlay", 0, ovfs.getOverlayOpt(imageId))
	if err != nil {
		return err
	}
	return nil
}
