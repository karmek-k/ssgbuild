package builder

import (
	"errors"
	"fmt"

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
	// check the build directory
	if err := CheckDir(cfg.BaseDir); err != nil {
		return fmt.Errorf("checking base dir failed: %s", err.Error())
	}

	// run the install command
	if utils.StringToCmd(cfg.InstallCmd).Run() != nil {
		return errors.New("failed running the install command")
	}

	// run the build command
	if utils.StringToCmd(cfg.BuildCmd).Run() != nil {
		return errors.New("failed running the install command")
	}

	// check the result directory
	if err := CheckDir(cfg.ResultDir); err != nil {
		return fmt.Errorf("checking result dir failed: %s", err.Error())
	} 

	return nil
}