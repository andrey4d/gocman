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
}

func (this *ContainerAttr) getRoot() string {
	pwd, err := os.Getwd()
	handlers.ErrorHandler(err, "PWD")
	return filepath.Join(pwd, this.Root)
}

func Child(cAtr ContainerAttr) {
	handlers.ErrorHandler(syscall.Chroot(cAtr.getRoot()), "change root")
	handlers.ErrorHandler(syscall.Chdir("/"), "change dir")

	cmd := exec.Command(cAtr.Command_name, cAtr.Arguments...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	MountProc()

	setHostname(cAtr.Container_name)
	fmt.Printf("Container: %d\n", os.Getpid())
	handlers.ErrorHandler(cmd.Run(), "run container")

	defer UmountProc()
}

func setHostname(hostname string) {
	handlers.ErrorHandler(syscall.Sethostname([]byte(hostname)), "set container name")
}
