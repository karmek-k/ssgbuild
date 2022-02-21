package phases

import (
	"errors"

	"github.com/karmek-k/ssgbuild/builder"
	"github.com/karmek-k/ssgbuild/utils"
)

type InstallCmdPhase struct {}

func (p InstallCmdPhase) Perform(b *builder.Build) error {
	b.Log.Debugw("running the install command",
		"cmd", b.Cfg.InstallCmd,
	)
	
	installOut, err := utils.StringToCmd(b.Cfg.InstallCmd).CombinedOutput()
	if err != nil {
		b.Log.Errorw("install command failed",
			"name", b.Cfg.Name,
			"out", string(installOut),
		)

		return errors.New("failed running the install command")
	}

	return nil
}