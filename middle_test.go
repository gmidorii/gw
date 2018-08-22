package gw_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/midorigreen/gw"
)

type testCmdImpl struct {
}

func (c testCmdImpl) Run(args []string, gotout, goterr io.Writer) error {
	fmt.Fprintf(gotout, "%v\n", args[0])
	fmt.Fprintf(goterr, "%v\n", args[1])
	return nil
}

func TestWrapEndEcho(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantout string
		wanterr string
	}{
		{
			name:    "health",
			in:      "hi",
			wantout: "gotout\nhi\n",
			wanterr: "goterr\n",
		},
	}

	c := testCmdImpl{}
	args := []string{"gotout", "goterr"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := gw.WrapEndEcho(tt.in)(c)
			gotout := bytes.Buffer{}
			goterr := bytes.Buffer{}

			err := cmd.Run(args, &gotout, &goterr)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if gotout.String() != tt.wantout {
				t.Errorf("gotout not equal got: %v, want: %v", gotout.String(), tt.wantout)
			}
			if goterr.String() != tt.wanterr {
				t.Errorf("goterr not equal got: %v, want: %v", goterr.String(), tt.wanterr)
			}
		})
	}
}

func TestWrapFirstEcho(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantout string
		wanterr string
	}{
		{
			name:    "health",
			in:      "start",
			wantout: "start\ngotout\n",
			wanterr: "goterr\n",
		},
	}

	c := testCmdImpl{}
	args := []string{"gotout", "goterr"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := gw.WrapFirstEcho(tt.in)(c)
			gotout := bytes.Buffer{}
			goterr := bytes.Buffer{}

			err := cmd.Run(args, &gotout, &goterr)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if gotout.String() != tt.wantout {
				t.Errorf("gotout not equal got: %v, want: %v", gotout.String(), tt.wantout)
			}
			if goterr.String() != tt.wanterr {
				t.Errorf("goterr not equal got: %v, want: %v", goterr.String(), tt.wanterr)
			}
		})
	}
}

func TestWrapTime(t *testing.T) {
	tests := []struct {
		name    string
		wantout string
		wanterr string
	}{
		{
			name:    "health",
			wantout: "gotout\ntime:",
			wanterr: "goterr\n",
		},
	}

	c := testCmdImpl{}
	args := []string{"gotout", "goterr"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := gw.WrapTime()(c)
			gotout := bytes.Buffer{}
			goterr := bytes.Buffer{}

			err := cmd.Run(args, &gotout, &goterr)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			// time could  not be fixed.
			if !strings.Contains(gotout.String(), tt.wantout) {
				t.Errorf("gotout not contains got: %v, want: %v", gotout.String(), tt.wantout)
			}
			if goterr.String() != tt.wanterr {
				t.Errorf("goterr not equal got: %v, want: %v", goterr.String(), tt.wanterr)
			}
		})
	}
}
