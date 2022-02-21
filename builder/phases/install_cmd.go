package phases

import (
	"errors"

	"github.com/karmek-k/ssgbuild/utils"
	"go.uber.org/zap"
)

func InstallCmdPhase(cfg *BuildConfig, log *zap.SugaredLogger) error {
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

	return nil
}