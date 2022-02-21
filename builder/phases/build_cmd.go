package phases

import (
	"errors"

	"github.com/karmek-k/ssgbuild/builder"
	"github.com/karmek-k/ssgbuild/utils"
)

type BuildCmdPhase struct {}

func (p BuildCmdPhase) Perform(b *builder.Build) error {
	b.Log.Debugw("running the build command",
		"cmd", b.Cfg.BuildCmd,
	)
	
	buildOut, err := utils.StringToCmd(b.Cfg.BuildCmd).CombinedOutput()
	if err != nil {
		b.Log.Errorw("build command failed",
			"name", b.Cfg.Name,
			"out", string(buildOut),
		)

		return errors.New("failed running the build command")
	}

	return nil
}