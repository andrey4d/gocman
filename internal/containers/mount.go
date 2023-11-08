package containers

import (
	"godman/internal/handlers"

	"syscall"
)

func MountProc() {
	handlers.ErrorHandler(syscall.Mount("proc", "/proc", "proc", 0, ""), "MountProc()")
}

func UmountProc() {
	handlers.ErrorHandler(syscall.Unmount("/proc", 0), "UmountProc()")
}
