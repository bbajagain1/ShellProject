package builtins

import (
	"os"
	"testing"
)

func TestMakeDirectory(t *testing.T) {
	// Create a temporary directory for testing.
	testDir := "testdir"
	// Ensure the directory is removed after the test finishes.
	defer os.RemoveAll(testDir)

	// Call the MakeDirectory function with the test directory as an argument.
	err := MakeDirectory(testDir)

	// Check if any error occurred.
	if err != nil {
		t.Errorf("MakeDirectory function returned an error: %v", err)
	}

	// Check if the test directory was created.
	_, err = os.Stat(testDir)
	if err != nil {
		if os.IsNotExist(err) {
			t.Errorf("MakeDirectory function did not create the test directory: %s", testDir)
		} else {
			t.Errorf("Error checking test directory: %v", err)
		}
	}
}
