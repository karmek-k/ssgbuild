package builder

import (
	"github.com/karmek-k/ssgbuild/builder/phases"
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

	// check & change to the base directory
	if err := phases.ChangeBaseDirPhase(cfg, log); err != nil {
		return err
	}

	// run the install command
	if err := phases.InstallCmdPhase(cfg, log); err != nil {
		return err
	}

	// run the build command
	if err := phases.BuildCmdPhase(cfg, log); err != nil {
		return err
	}

	// check the result directory
	if err := phases.CheckResultPhase(cfg, log); err != nil {
		return err
	}

	log.Infow("build succeeded",
		"name", cfg.Name,
	)

	return nil
}