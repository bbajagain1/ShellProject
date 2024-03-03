package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidArgCount = errors.New("invalid argument count")
)

func PrintWorkingDirectory(args ...string) error {
	// Check the number of arguments.
	if len(args) > 0 {
		return fmt.Errorf("%w: expected zero arguments", ErrInvalidArgCount)
	}

	// Get the current working directory.
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Print the current working directory.
	fmt.Println(wd)
	return nil
}
