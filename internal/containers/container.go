package containers

import (
	"fmt"
	"godman/internal/helpers"
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

	helpers.ErrorHelperPanicWithMessage(MountOvfs(cAtr.OvfsRoot), "mount overlay")

	MountProc(GetAbsPath(cAtr.OvfsRoot.Target))

	helpers.ErrorHelperPanicWithMessage(MountRoot(GetAbsPath(cAtr.OvfsRoot.Target)), "pivot root")

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
