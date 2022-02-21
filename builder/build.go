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

	// run all phases in order
	for _, phase := range phases {
		if err := phase.Perform(b); err != nil {
			return err
		}
	}

	b.Log.Infow("build succeeded",
		"name", b.Cfg.Name,
	)

	return nil
}