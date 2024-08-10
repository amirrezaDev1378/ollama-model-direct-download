package app

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
)

func HasElevatedPermissions() (bool, error) {

	switch runtime.GOOS {
	case "windows":
		_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
		return err == nil, nil
	case "linux", "darwin":
		currentUser, err := user.Current()
		if err != nil {
			return false, err
		}
		return currentUser.Uid == "0", nil
	default:
		return false, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}
