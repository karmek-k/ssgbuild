package utils

import (
	"fmt"
	"os"
)

// CheckDir checks if dir exists and is writable
func CheckDir(dir string) error {
	stat, err := os.Stat(dir)
	if err != nil {
		return err
	}

	// is dir a directory?
	if !stat.IsDir() {
		return fmt.Errorf("not a directory: %s", dir)
	}

	return nil
}
