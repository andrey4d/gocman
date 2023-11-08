package starter

import (
	"fmt"
	"godman/internal/handlers"
	"os"
	"os/exec"
	"syscall"
)

func Run(args []string) {
	fmt.Println("Hello godman!")

	command := append([]string{"child"}, args[2:]...)

	cmd := exec.Command("/proc/self/exe", command...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	cmd.SysProcAttr.Cloneflags = syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS

	fmt.Printf("Starter: %d\n", os.Getpid())
	handlers.ErrorHandler(cmd.Run(), "Run() ")
}
