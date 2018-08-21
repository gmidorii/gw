package gw

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func Chain(out CmdMiddleware, mid ...CmdMiddleware) CmdMiddleware {
	return func(c Cmder) Cmder {
		fidx := len(mid) - 1
		for i := range mid {
			c = mid[fidx-i](c)
		}
		return out(c)
	}
}

func WrapEndEcho(s string) CmdMiddleware {
	return func(c Cmder) Cmder {
		fn := func(args []string, stdout, stderr io.Writer) error {
			defer fmt.Printf("%v\n", s)
			return c.Run(args, stdout, stderr)
		}
		return CmdFunc(fn)
	}
}

func WrapFirstEcho(s string) CmdMiddleware {
	return func(c Cmder) Cmder {
		fn := func(args []string, stdout, stderr io.Writer) error {
			fmt.Printf("%v\n", s)
			return c.Run(args, stdout, stderr)
		}
		return CmdFunc(fn)
	}
}

func WrapTime() CmdMiddleware {
	return func(c Cmder) Cmder {
		fn := func(args []string, stdout, stderr io.Writer) error {
			s := time.Now()
			defer fmt.Printf("time:%v \n", time.Now().Sub(s))
			return c.Run(args, stdout, stderr)
		}
		return CmdFunc(fn)
	}
}

func WrapNotify(notifier Notifier, dest string) CmdMiddleware {
	return func(c Cmder) Cmder {
		fn := func(args []string, stdout, stderr io.Writer) error {
			err := c.Run(args, stdout, stderr)
			title := strings.Join(args, "\\s")
			if err != nil {
				return notifier.Send(title, dest, fmt.Sprintln(err))
			}
			return notifier.Send(title, dest, fmt.Sprintln(err))
		}
		return CmdFunc(fn)
	}
}
