package builtins

import (
	"bytes"
	"testing"
)

func TestListDirectory(t *testing.T) {
	// Create a buffer to capture the output.
	var buf bytes.Buffer

	// Call the ListDirectory function.
	err := ListDirectory()

	// Check if any error occurred.
	if err != nil {
		t.Errorf("ListDirectory function returned an error: %v", err)
	}
}
