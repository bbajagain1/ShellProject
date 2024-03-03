package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidArgCount = errors.New("invalid argument count")
)

func MakeDirectory(args ...string) error {
	// Check the number of arguments.
	if len(args) == 0 {
		return fmt.Errorf("%w: missing operand", ErrInvalidArgCount)
	}

	// Create directories with the given names.
	for _, dir := range args {
		err := os.Mkdir(dir, 0755) // 0755 is the default permission mode for mkdir
		if err != nil {
			return err
		}
	}

	return nil
}
