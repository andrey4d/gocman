/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package containers

import (
	"fmt"
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

type ContainerAttr struct {
	Command_name   string
	Arguments      []string
	Container_name string
	OvfsRoot       *OvfsMountCfg
}

func Container(cAtr ContainerAttr) {
	fmt.Printf("change root to %s\n", helpers.GetAbsPath(cAtr.OvfsRoot.Target))

	helpers.ErrorHelperPanicWithMessage(MountOvfs(cAtr.OvfsRoot), "mount overlay")

	MountProc(helpers.GetAbsPath(cAtr.OvfsRoot.Target))

	helpers.ErrorHelperPanicWithMessage(MountRoot(helpers.GetAbsPath(cAtr.OvfsRoot.Target)), "pivot root")

	helpers.ErrorHelperPanicWithMessage(syscall.Chdir("/"), "change dir")

	cmd := exec.Command(cAtr.Command_name, cAtr.Arguments...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{"PATH=/bin:/sbin:/usr/bin:/usr/sbin"}

	setHostname(cAtr.Container_name)

	fmt.Printf("Container PID: %d\n", os.Getpid())
	helpers.ErrorHelperPanicWithMessage(cmd.Run(), "run container")

	helpers.ErrorHelperPanicWithMessage(UmountProc(), "umount /proc")
}

func setHostname(hostname string) {
	helpers.ErrorHelperPanicWithMessage(syscall.Sethostname([]byte(hostname)), "set container name")
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
		helpers.ErrorHelperPanicWithMessage(o.MkLower(), "make overlay lower")
		// STUB
		helpers.ErrorHelperPanicWithMessage(helpers.Untar("alpine.tar", o.Lowerdir[0]), "untar image")
		// STUB
		wg.Done()
	}()

	for _, point := range mounts {
		go func(p string) {
			helpers.ErrorHelperPanicWithMessage(helpers.MakeDirAllIfNotExists(p, o.Permeations), fmt.Sprintf("make overlay %s\n", p))
			wg.Done()
		}(point)
	}
	wg.Wait()
}
