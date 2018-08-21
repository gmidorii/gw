package gw

import (
	"io"
)

type Cmder interface {
	Run(args []string, stdout, stderr io.Writer) error
}

type CmdMiddleware func(Cmder) Cmder

type CmdFunc func(args []string, stdout, stderr io.Writer) error

func (c CmdFunc) Run(args []string, stdout, stderr io.Writer) error {
	return c(args, stdout, stderr)
}
