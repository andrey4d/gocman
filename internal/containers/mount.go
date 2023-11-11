package containers

import (
	"fmt"
	"io/fs"
	"sync"

	"godman/internal/helpers"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func MountProc(new_root string) {
	target := filepath.Join(new_root, "/proc")
	helpers.ErrorHelperPanicWithMessage(syscall.Mount("proc", target, "proc", 0, ""), "MountProc()")
}

func UmountProc() error {
	if err := syscall.Unmount("/proc", 0); err != nil {
		return err
	}
	return nil
}

func MountRoot(new_root string) error {

	old_root := filepath.Join(new_root, "/.pivot_root")
	fmt.Println(new_root)
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

type OvfsMountCfg struct {
	Lowerdir    []string // overlay/l/<ZLR4NWYDXWB5LCOBDH7WAGVYDI> --> storage/overlay/<id>/diff
	Upperdir    string   // storage/overlay/<id>/diff
	Workdir     string   // storage/overlay/<id>/work
	Target      string   // storage/overlay/<id>/merged
	Permeations fs.FileMode
	SELebel     string
}

func (o *OvfsMountCfg) OvfsOpt() string {

	var lowerdir []string
	for _, s := range o.Lowerdir {
		lowerdir = append(lowerdir, GetAbsPath(s))
	}
	opts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", strings.Join(lowerdir, ":"), GetAbsPath(o.Upperdir), GetAbsPath(o.Workdir))
	return opts
}

func (o *OvfsMountCfg) makeOvfsDir(dir string, chmod fs.FileMode) error {
	if _, err := os.Stat(dir); os.IsExist(err) {
		return nil
	}

	if err := os.MkdirAll(dir, chmod); err != nil {
		return err
	}
	return nil
}

func (o *OvfsMountCfg) MkUpper() error {
	return o.makeOvfsDir(o.Upperdir, o.Permeations)
}

func (o *OvfsMountCfg) MkWork() error {
	return o.makeOvfsDir(o.Workdir, o.Permeations)
}

func (o *OvfsMountCfg) MkTarget() error {
	return o.makeOvfsDir(o.Target, o.Permeations)
}

func (o *OvfsMountCfg) MkLower() error {
	for _, v := range o.Lowerdir {
		if err := o.makeOvfsDir(v, o.Permeations); err != nil {
			return err
		}
	}
	return nil
}

func (o *OvfsMountCfg) MkAll() error {
	if err := o.MkUpper(); err != nil {
		return err
	}
	if err := o.MkWork(); err != nil {
		return err
	}
	if err := o.MkTarget(); err != nil {
		return err
	}
	if err := o.MkLower(); err != nil {
		return err
	}
	return nil
}

func MountOvfs(ovfs *OvfsMountCfg) error {

	mounts := []string{
		ovfs.Workdir,
		ovfs.Upperdir,
		ovfs.Target,
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(mounts) + 1)

	go func() {
		helpers.ErrorHelperPanicWithMessage(ovfs.MkLower(), "make overlay lower")
		// STUB
		helpers.ErrorHelperPanicWithMessage(helpers.Untar("alpine.tar", ovfs.Lowerdir[0]), "untar image")
		// STUB
		wg.Done()
	}()

	for _, point := range mounts {
		go func(p string) {
			helpers.ErrorHelperPanicWithMessage(ovfs.makeOvfsDir(p, ovfs.Permeations), fmt.Sprintf("make overlay %s\n", p))
			wg.Done()
		}(point)
	}
	wg.Wait()

	target := GetAbsPath(ovfs.Target)
	opts := ovfs.OvfsOpt()
	fmt.Print("mount overlay fs root ...\n")
	fmt.Println(opts)

	err := syscall.Mount("overlay", target, "overlay", 0, opts)
	if err != nil {
		return err
	}

	return nil
}
