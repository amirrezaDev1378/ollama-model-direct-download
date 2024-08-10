package app

import (
	"fmt"
	"os"
	"slices"
)

func VerifyDownloadedModel(modelName string, downloadedModelPath string) error {
	fileInfo, err := os.Stat(downloadedModelPath)
	if err != nil {
		return fmt.Errorf("error checking target folder ( %v ) : %v", downloadedModelPath, err.Error())
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("downloaded model path is not a directory")
	}

	files, err := os.ReadDir(downloadedModelPath)
	if err != nil {
		return fmt.Errorf("error reading downloaded model directory: %v", err.Error())
	}
	if !(len(files) > 2) {
		return fmt.Errorf("downloaded model directory does not contain required files")
	}

	containsManifest := slices.ContainsFunc(files, func(file os.DirEntry) bool {
		return file.Name() == "manifest"
	})

	if !containsManifest {
		return fmt.Errorf("downloaded model directory does not contain the manifest file")
	}

	return nil
}
