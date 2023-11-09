package containers

import (
	"fmt"
	"godman/internal/handlers"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type ContainerAttr struct {
	Command_name   string
	Arguments      []string
	Container_name string
	Root           string
	OvfsRoot       *OvfsMountCfg
}

func (c *ContainerAttr) getRoot() string {

	if filepath.IsAbs(c.Root) {
		return c.Root
	}

	pwd, err := os.Getwd()
	handlers.ErrorHandlerPanicWithMessage(err, "PWD")
	return filepath.Join(pwd, c.Root)

}

func Container(cAtr ContainerAttr) {
	fmt.Printf("change root to %s\n", cAtr.getRoot())

	handlers.ErrorHandlerPanicWithMessage(MountOvfs(cAtr.OvfsRoot), "mount overlay")
	MountProc(cAtr.getRoot())
	handlers.ErrorHandlerPanicWithMessage(MountRoot(cAtr.getRoot()), "pivot root")

	handlers.ErrorHandlerPanicWithMessage(syscall.Chdir("/"), "change dir")

	cmd := exec.Command(cAtr.Command_name, cAtr.Arguments...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{"PATH=/bin:/sbin:/usr/bin:/usr/sbin"}

	setHostname(cAtr.Container_name)
	fmt.Printf("Container PID: %d\n", os.Getpid())
	handlers.ErrorHandlerPanicWithMessage(cmd.Run(), "run container")

	defer UmountProc()
}

func setHostname(hostname string) {
	handlers.ErrorHandlerPanicWithMessage(syscall.Sethostname([]byte(hostname)), "set container name")
}
