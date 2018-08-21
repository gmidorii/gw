package main

import (
	"log"
	"os"
)

func run(args []string) error {
	cmd := NewCmd()
	return cmd.Run(args, os.Stdout, os.Stderr)
}

func main() {
	args := os.Args
	if err := run(args[1:]); err != nil {
		log.Fatalf("failed wrap: %v", err)
	}
}
