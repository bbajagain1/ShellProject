package builtins

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	// Create a buffer to capture the output.
	var buf bytes.Buffer
	// Arguments to be passed to the Echo function.
	args := []string{"Hello", "World"}

	// Call the Echo function with the provided arguments.
	err := Echo(args...)

	// Check if any error occurred.
	if err != nil {
		t.Errorf("Echo function returned an error: %v", err)
	}

	// Check if the output matches the expected output.
	expectedOutput := "Hello World\n"
	if buf.String() != expectedOutput {
		t.Errorf("Echo function output does not match. Expected: %q, Got: %q", expectedOutput, buf.String())
	}
}
