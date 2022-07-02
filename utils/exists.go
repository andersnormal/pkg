package utils

import (
	"os"
)

// FileExists is testing if a file exists at a given path.
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false, err
	}

	return true, nil
}
