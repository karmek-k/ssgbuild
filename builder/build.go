package builder

import (
	"errors"
	"os"
)

func Build(baseDir, command, resultDir string) error {
	os.Chdir(baseDir)
	
	if _, err := os.Getwd(); err != nil {
		return errors.New("failed changing to the base directory")
	}

	// err := exec.Command(command).Run()
	// if err != nil {
	// 	return errors.New("failed running the build command")
	// }

	return nil
}