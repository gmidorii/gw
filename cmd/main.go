package main

import (
	"errors"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/midorigreen/gw"
)

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

func run(args []string) error {
	var cmd gw.Cmder = cmdImpl{}
	cmd = gw.Chain(
		gw.WrapTime(),
		gw.WrapFirstEcho("=== START ==="),
		gw.WrapSlack(os.Getenv("SLACK_TOKEN"), "send test", os.Getenv("SLACK_CHANNEL")),
		gw.WrapEndEcho("=== END ==="),
	)(cmd)
	return cmd.Run(args, os.Stdout, os.Stderr)
}

func main() {
	args := os.Args
	if err := run(args[1:]); err != nil {
		log.Fatalf("failed wrap: %v", err)
	}
}
