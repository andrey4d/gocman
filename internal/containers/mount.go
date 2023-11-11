package containers

import (
	"fmt"
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
	Lowerdir []string
	Upperdir string
	Workdir  string
	Target   string
	SELebel  string
}

func (o *OvfsMountCfg) OvfsOpt() string {

	var lowerdir []string
	for _, s := range o.Lowerdir {
		lowerdir = append(lowerdir, helpers.GetAbsPath(s))
	}

	opts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", strings.Join(lowerdir, ":"), helpers.GetAbsPath(o.Upperdir), helpers.GetAbsPath(o.Workdir))

	return opts
}

func MountOvfs(ovfs *OvfsMountCfg) error {

	target := helpers.GetAbsPath(ovfs.Target)
	opts := ovfs.OvfsOpt()
	fmt.Print("mount overlay fs root ...\n")
	fmt.Println(opts)

	err := syscall.Mount("overlay", target, "overlay", 0, opts)
	if err != nil {
		return err
	}

	return nil
}
