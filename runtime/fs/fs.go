package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

func FileExists(path string) (bool, error) {
	expandedPath := os.ExpandEnv(path)

	_, err := os.Stat(expandedPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func RootMkdirP(dirPath string) error {
	expandedPath := os.ExpandEnv(dirPath)
	absPath, err := filepath.Abs(expandedPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	err = os.MkdirAll(absPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	return nil
}
