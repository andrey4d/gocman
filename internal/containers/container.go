package containers

import (
	"fmt"
	"godman/internal/handlers"
	"os"
	"os/exec"
	"syscall"
)

type ContainerAttr struct {
	Command_name   string
	Arguments      []string
	Container_name string
	OvfsRoot       *OvfsMountCfg
}

func Container(cAtr ContainerAttr) {
	fmt.Printf("change root to %s\n", GetAbsPath(cAtr.OvfsRoot.Target))

	handlers.ErrorHandlerPanicWithMessage(MountOvfs(cAtr.OvfsRoot), "mount overlay")

	MountProc(GetAbsPath(cAtr.OvfsRoot.Target))

	handlers.ErrorHandlerPanicWithMessage(MountRoot(GetAbsPath(cAtr.OvfsRoot.Target)), "pivot root")

	handlers.ErrorHandlerPanicWithMessage(syscall.Chdir("/"), "change dir")

	cmd := exec.Command(cAtr.Command_name, cAtr.Arguments...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{"PATH=/bin:/sbin:/usr/bin:/usr/sbin"}

	setHostname(cAtr.Container_name)
	fmt.Printf("Container PID: %d\n", os.Getpid())
	handlers.ErrorHandlerPanicWithMessage(cmd.Run(), "run container")

	handlers.ErrorHandlerPanicWithMessage(UmountProc(), "umount /proc")
}

func setHostname(hostname string) {
	handlers.ErrorHandlerPanicWithMessage(syscall.Sethostname([]byte(hostname)), "set container name")
}
