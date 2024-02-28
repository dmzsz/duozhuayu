package helpers_test

import (
	"os"
	"testing"

	"github.com/dmzsz/duozhuayu/pkg/helpers"
)

func TestFileExist(t *testing.T) {
	// create temporary file
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// test with existing file
	if !helpers.FileExist(tempFile.Name()) {
		t.Error("expected existing file to return true, but got false")
	}

	// test with non-existing file
	if helpers.FileExist("non-existing-file") {
		t.Error("expected non-existing file to return false, but got true")
	}
}
