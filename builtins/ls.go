package builtins

import (
	"errors"
	"os"
	"os/exec"
)

var (
	ErrInvalidArgCount = errors.New("invalid argument count")
)

func ListDirectory(args ...string) error {
	// Check the number of arguments.
	if len(args) > 0 {
		return ErrInvalidArgCount
	}

	// Execute the `ls` command.
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
