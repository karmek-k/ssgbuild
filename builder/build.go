package builder

import (
	"errors"
	"fmt"
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
	// check the build directory
	if err := CheckDir(cfg.BaseDir); err != nil {
		return fmt.Errorf("checking base dir failed: %s", err.Error())
	}

	// change to the build directory
	if os.Chdir(cfg.BaseDir) != nil {
		return errors.New("could not change the base directory")
	}

	// run the install command
	installOut, err := utils.StringToCmd(cfg.InstallCmd).CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"failed running the install command - output:\n%s",
			string(installOut),
		)
	}

	// run the build command
	buildOut, err := utils.StringToCmd(cfg.BuildCmd).CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"failed running the build command - output:\n%s",
			string(buildOut),
		)
	}

	// check the result directory
	if err := CheckDir(cfg.ResultDir); err != nil {
		return fmt.Errorf("checking result dir failed: %s", err.Error())
	}

	return nil
}