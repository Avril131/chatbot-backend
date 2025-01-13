package utils

import "os"

// PathExists checks if the given path exists or not.
// It returns a boolean indicating existence and an error if any occurs.
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
