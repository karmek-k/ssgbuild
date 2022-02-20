package builder

import (
	"errors"
	"os"
	"os/exec"
)

// Build builds the website with given parameters
// and returns an error if there is any
func Build(baseDir, installCmd, buildCmd, resultDir string) error {
	// change to the build directory
	os.Chdir(baseDir)
	if _, err := os.Getwd(); err != nil {
		return errors.New("failed changing to the base directory")
	}

	// run the install command
	if exec.Command(installCmd).Run() != nil {
		return errors.New("failed running the install command")
	}

	// run the build command
	if exec.Command(buildCmd).Run() != nil {
		return errors.New("failed running the install command")
	}

	return nil
}