package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type cmd struct {
}

func NewCmd() *cmd {
	return &cmd{}
}

func (c *cmd) Run(args []string, stdout, stderr io.Writer) error {
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
