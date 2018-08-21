package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("must arguments >= 1")
	}
	fmt.Println(args)
	var cmd *exec.Cmd
	if len(args) == 1 {
		cmd = exec.Command(args[0])
	} else {
		cmd = exec.Command(args[0], args[1:]...)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	args := os.Args
	if err := run(args[1:]); err != nil {
		log.Fatalf("failed wrap: %v", err)
	}
}
