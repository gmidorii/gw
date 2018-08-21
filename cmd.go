package main

import (
	"errors"
	"io"
	"os"
	"os/exec"
)

type Cmder interface {
	Run(args []string, stdout, stderr io.Writer) error
}

type cmdFunc func(args []string, stdout, stderr io.Writer) error

func (c cmdFunc) Run(args []string, stdout, stderr io.Writer) error {
	return c(args, stdout, stderr)
}

type cmdMiddleware func(Cmder) Cmder

type cmdImpl struct{}

func (c cmdImpl) Run(args []string, stdout, stderr io.Writer) error {
	if len(args) == 0 {
		return errors.New("must arguments >= 1")
	}
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
