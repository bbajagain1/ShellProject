package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidArgCount = errors.New("invalid argument count")
)

func RemoveDirectory(args ...string) error {
	// Check the number of arguments.
	if len(args) == 0 {
		return fmt.Errorf("%w: missing operand", ErrInvalidArgCount)
	}

	// Remove directories.
	for _, dir := range args {
		err := os.Remove(dir)
		if err != nil {
			return err
		}
	}

	return nil
}
