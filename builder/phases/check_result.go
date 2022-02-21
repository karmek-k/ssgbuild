package phases

import (
	"fmt"

	"github.com/karmek-k/ssgbuild/utils"
	"go.uber.org/zap"
)

func CheckResultPhase(cfg *BuildConfig, log *zap.SugaredLogger) error {
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

	return nil
}