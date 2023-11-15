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
	"os/exec"
	"strings"
	"sync"
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

func Container(args []string) {

	cId := helpers.CreateContainerID(16)

	containerBaseDir := fmt.Sprintf("%s/storage/overlay/%s", helpers.GetAbsPath(config.Config.GetContainersPath()), cId)

	var ovfsMountCfg = OvfsMountCfg{
		Lowerdir:    []string{fmt.Sprintf("%s/l", containerBaseDir)}, // overlay/l/<ZLR4NWYDXWB5LCOBDH7WAGVYDI> --> storage/overlay/<id>/diff
		Upperdir:    fmt.Sprintf("%s/diff", containerBaseDir),        // storage/overlay/<id>/diff
		Workdir:     fmt.Sprintf("%s/work", containerBaseDir),        // storage/overlay/<id>/work
		Target:      fmt.Sprintf("%s/merged", containerBaseDir),      // storage/overlay/<id>/merged
		Permeations: config.Config.GetPermissions(),
		SELebel:     "",
	}

	command_name := args[0]
	arguments := args[1:]
	container_name := cId

	fmt.Printf("change root to %s\n", helpers.GetAbsPath(ovfsMountCfg.Target))

	helpers.CheckError(MountOvfs(&ovfsMountCfg), "mount overlay")

	MountProc(helpers.GetAbsPath(ovfsMountCfg.Target))

	helpers.CheckError(MountRoot(helpers.GetAbsPath(ovfsMountCfg.Target)), "pivot root")

	helpers.CheckError(syscall.Chdir("/"), "change dir")

	cmd := exec.Command(command_name, arguments...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{"PATH=/bin:/sbin:/usr/bin:/usr/sbin"}

	setHostname(container_name)

	fmt.Printf("Container PID: %d\n", os.Getpid())
	helpers.CheckError(cmd.Run(), "run container")

	helpers.CheckError(UmountProc(), "umount /proc")
}

func setHostname(hostname string) {
	helpers.CheckError(syscall.Sethostname([]byte(hostname)), "set container name")
}

func (o *OvfsMountCfg) MkUpper() error {
	return helpers.MakeDirAllIfNotExists(o.Upperdir, o.Permeations)
}

func (o *OvfsMountCfg) MkWork() error {
	return helpers.MakeDirAllIfNotExists(o.Workdir, o.Permeations)
}

func (o *OvfsMountCfg) MkTarget() error {
	return helpers.MakeDirAllIfNotExists(o.Target, o.Permeations)
}

func (o *OvfsMountCfg) MkLower() error {
	for _, v := range o.Lowerdir {
		if err := helpers.MakeDirAllIfNotExists(v, o.Permeations); err != nil {
			return err
		}
	}
	return nil
}

func (o *OvfsMountCfg) OvfsOpt() string {

	var lowerdir []string
	for _, s := range o.Lowerdir {
		lowerdir = append(lowerdir, helpers.GetAbsPath(s))
	}
	opts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", strings.Join(lowerdir, ":"), helpers.GetAbsPath(o.Upperdir), helpers.GetAbsPath(o.Workdir))
	return opts
}

func (o *OvfsMountCfg) createDirectoryStructure() {
	mounts := []string{
		o.Workdir,
		o.Upperdir,
		o.Target,
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(mounts) + 1)

	go func() {
		helpers.CheckError(o.MkLower(), "make overlay lower")
		// STUB
		helpers.CheckError(helpers.Untar("alpine.tar", o.Lowerdir[0]), "untar image")
		// STUB
		wg.Done()
	}()

	for _, point := range mounts {
		go func(p string) {
			helpers.CheckError(helpers.MakeDirAllIfNotExists(p, o.Permeations), fmt.Sprintf("make overlay %s\n", p))
			wg.Done()
		}(point)
	}
	wg.Wait()
}
