package phases

import (
	"errors"

	"github.com/karmek-k/ssgbuild/utils"
	"go.uber.org/zap"
)

func BuildCmdPhase(cfg *BuildConfig, log *zap.SugaredLogger) error {
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

	return nil
}