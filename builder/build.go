package builder

import (
	"errors"
	"fmt"
	"os"

	"github.com/karmek-k/ssgbuild/utils"
	"go.uber.org/zap"
)

// BuildConfig struct defines how the site should be built
type BuildConfig struct {
	Name string
	BaseDir string
	InstallCmd string
	BuildCmd string
	ResultDir string
}

// Build builds the website with given parameters
// and returns an error if there is any
func Build(cfg *BuildConfig, log *zap.SugaredLogger) error {
	log.Infow("starting a build",
		"name", cfg.Name,
	)

	// check & chango to the base directory
	log.Debugw("changing to the base directory",
		"dir", cfg.BaseDir,
	)
	if err := utils.CheckDir(cfg.BaseDir); err != nil {
		log.Errorw("checking base dir failed",
			"name", cfg.Name,
			"dir", cfg.BaseDir,
			"error", err.Error(),
		)

		return fmt.Errorf("checking base dir failed: %s", err.Error())
	}

	if os.Chdir(cfg.BaseDir) != nil {
		log.Errorw("changing to base dir failed",
			"name", cfg.Name,
			"dir", cfg.BaseDir,
		)

		return errors.New("could not change the base directory")
	}

	// run the install command
	log.Debugw("running the install command",
		"cmd", cfg.InstallCmd,
	)
	installOut, err := utils.StringToCmd(cfg.InstallCmd).CombinedOutput()
	if err != nil {
		log.Errorw("install command failed",
			"name", cfg.Name,
			"out", string(installOut),
		)

		return errors.New("failed running the install command")
	}

	// run the build command
	log.Debugw("running the build command",
		"cmd", cfg.BuildCmd,
	)
	buildOut, err := utils.StringToCmd(cfg.BuildCmd).CombinedOutput()
	if err != nil {
		log.Errorw("build command failed",
			"name", cfg.Name,
			"out", string(buildOut),
		)

		return errors.New("failed running the build command")
	}

	// check the result directory
	log.Debugw("checking the result dir",
		"dir", cfg.ResultDir,
	)
	if err := utils.CheckDir(cfg.ResultDir); err != nil {
		log.Errorw("checking result dir failed",
			"name", cfg.Name,
			"dir", cfg.ResultDir,
		)

		return fmt.Errorf("checking result dir failed: %s", err.Error())
	}

	log.Infow("build succeeded",
		"name", cfg.Name,
	)

	return nil
}