package app

import (
	"os"
	"testing"
)

func TestVerifyDownloadedModel(t *testing.T) {
	tempDir := t.TempDir()
	defer os.Remove(tempDir)

	manifestPath := tempDir + "/manifest"
	manifestFile, err := os.Create(manifestPath)
	defer manifestFile.Close()
	if err != nil {
		t.Fatalf("Error creating manifest file: %v", err)
	}

	dummyFile1 := tempDir + "/sha256dummy1"
	tempFile, err := os.Create(dummyFile1)
	defer tempFile.Close()
	if err != nil {
		t.Fatalf("Error creating dummyFile1 : %v", err)
	}

	dummyFile2 := tempDir + "/sha256dummy2"
	tempFile, err = os.Create(dummyFile2)
	defer tempFile.Close()

	if err != nil {
		t.Fatalf("Error creating dummyFile2 : %v", err)
	}

	err = VerifyDownloadedModel("myModel", tempDir)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	err = VerifyDownloadedModel("myModel", manifestPath)
	if err == nil {
		t.Error("Expected an error for non-directory path, but got none")
	}
	manifestFile.Close()
	err = os.Remove(manifestPath)
	if err != nil {
		t.Fatalf("Error removing manifest file: %v", err)
	}
	err = VerifyDownloadedModel("myModel", tempDir)
	if err == nil {
		t.Error("Expected an error for missing manifest file, but got none")
	}
}
