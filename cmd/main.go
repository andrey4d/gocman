package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s run <cmd>\n", os.Args[0])
		os.Exit(0)
	}

	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help: use run <cmd>")
	}
}

func run() {
	fmt.Println("Hello godman!")

	command := append([]string{"child"}, os.Args[2:]...)

	cmd := exec.Command("/proc/self/exe", command...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	cmd.SysProcAttr.Cloneflags = syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID

	fmt.Printf("Starter: %d\n", os.Getpid())
	errorHandler(cmd.Run())

}

func child() {

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Container: %d\n", os.Getpid())
	errorHandler(cmd.Run())

}

func errorHandler(err error) {
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

}
