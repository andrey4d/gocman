package containers

import (
	"fmt"
	"godman/internal/handlers"
	"os"
	"path/filepath"
	"strings"

	"syscall"
)

func MountProc(new_root string) {
	target := filepath.Join(new_root, "/proc")
	handlers.ErrorHandlerPanicWithMessage(syscall.Mount("proc", target, "proc", 0, ""), "MountProc()")
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
		handlers.ErrorHandlerLog("Bind newroot")
		return err
	}

	if err := os.MkdirAll(old_root, 0700); err != nil {
		handlers.ErrorHandlerLog("make dirs " + old_root)
		return err
	}

	if err := syscall.PivotRoot(new_root, old_root); err != nil {
		handlers.ErrorHandlerLog("syscall.PivotRoot")
		return err
	}

	if err := syscall.Chdir("/"); err != nil {
		handlers.ErrorHandlerLog("change dir to /")
		return err
	}

	if err := syscall.Unmount("/.pivot_root", syscall.MNT_DETACH); err != nil {
		handlers.ErrorHandlerLog("umount " + old_root)
		return err
	}

	if err := os.RemoveAll("/.pivot_root"); err != nil {
		handlers.ErrorHandlerLog("remove " + old_root)
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
		lowerdir = append(lowerdir, GetAbsPath(s))
	}

	opts := fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", strings.Join(lowerdir, ":"), GetAbsPath(o.Upperdir), GetAbsPath(o.Workdir))

	return opts
}

func GetAbsPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	pwd, err := os.Getwd()
	handlers.ErrorHandlerPanicWithMessage(err, "PWD in GetAbsPath")
	return filepath.Join(pwd, path)
}

func MountOvfs(ovfs *OvfsMountCfg) error {

	target := GetAbsPath(ovfs.Target)
	opts := ovfs.OvfsOpt()
	fmt.Println(opts)
	fmt.Println(target)

	err := syscall.Mount("overlay", target, "overlay", 0, opts)
	if err != nil {
		return err
	}

	return nil
}
