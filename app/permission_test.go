package app

import (
	"testing"
)

func TestHasElevatedPermissions(t *testing.T) {

	// Check if the current process has elevated permissions
	_, err := HasElevatedPermissions()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
