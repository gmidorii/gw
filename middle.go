package main

import (
	"fmt"
	"io"
)

func WrapEcho(s string) cmdMiddleware {
	return func(c Cmder) Cmder {
		fn := func(args []string, stdout, stderr io.Writer) error {
			defer fmt.Printf("%v\n", s)
			return c.Run(args, stdout, stderr)
		}
		return cmdFunc(fn)
	}
}
