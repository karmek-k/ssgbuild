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

func buildConfigToMap(c *BuildConfig) phases.PhaseArgs {
	return phases.PhaseArgs{
		"Name": c.Name,
		"BaseDir": c.BaseDir,
		"InstallCmd": c.InstallCmd,
		"BuildCmd": c.BuildCmd,
		"ResultDir": c.ResultDir,
	}
}

type Build struct {
	Cfg *BuildConfig
	Log *zap.SugaredLogger
}

// Start builds the website with given parameters
// and returns an error if there is any
func (b *Build) Start(phases []phases.Phase) error {
	b.Log.Infow("starting a build",
		"name", b.Cfg.Name,
	)

	// convert the BuildConfig object to map[string]string
	// in order to prevent circular importing
	// TODO: maybe there is a better way of achieving this?
	args := buildConfigToMap(b.Cfg)

	// run all phases in order
	for _, phase := range phases {
		b.Log.Debugw("running build phase",
			"build", b.Cfg.Name,
			"phase", phase.GetName(),
		)

		if err := phase.Perform(args); err != nil {
			b.Log.Errorw("build failed",
				"build", b.Cfg.Name,
				"phase", phase.GetName(),
				"error", err.Error(),
			)

			return err
		}
	}

	b.Log.Infow("build succeeded",
		"name", b.Cfg.Name,
	)

	return nil
}

