package builder

import (
	"errors"
	"os"
	"os/exec"
)

// BuildConfig struct defines how the site should be built
type BuildConfig struct {
	BaseDir string
	InstallCmd string
	BuildCmd string
	ResultDir string
}

// Build builds the website with given parameters
// and returns an error if there is any
func Build(cfg BuildConfig) error {
	// change to the build directory
	os.Chdir(cfg.BaseDir)
	if _, err := os.Getwd(); err != nil {
		return errors.New("failed changing to the base directory")
	}

	// run the install command
	if exec.Command(cfg.InstallCmd).Run() != nil {
		return errors.New("failed running the install command")
	}

	// run the build command
	if exec.Command(cfg.BuildCmd).Run() != nil {
		return errors.New("failed running the install command")
	}

	return nil
}