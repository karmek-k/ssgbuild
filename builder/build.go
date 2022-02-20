package builder

import (
	"errors"
	"os"

	"github.com/karmek-k/ssgbuild/utils"
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
func Build(cfg *BuildConfig) error {
	// change to the build directory
	os.Chdir(cfg.BaseDir)
	if _, err := os.Getwd(); err != nil {
		return errors.New("failed changing to the base directory")
	}

	// run the install command
	if utils.StringToCmd(cfg.InstallCmd).Run() != nil {
		return errors.New("failed running the install command")
	}

	// run the build command
	if utils.StringToCmd(cfg.BuildCmd).Run() != nil {
		return errors.New("failed running the install command")
	}

	return nil
}