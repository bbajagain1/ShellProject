package builtins

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidArgCount = errors.New("invalid argument count")
)

// Echo prints the provided arguments to standard output.
func Echo(args ...string) error {
	// Check the number of arguments.
	if len(args) == 0 {
		return fmt.Errorf("%w: missing operand", ErrInvalidArgCount)
	}

	// Print the arguments separated by spaces.
	fmt.Println(strings.Join(args, " "))

	return nil
}
