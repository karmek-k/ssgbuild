package phases

import (
	"errors"
	"fmt"
	"os"

	"github.com/karmek-k/ssgbuild/utils"
	"go.uber.org/zap"
)

func ChangeBaseDirPhase(cfg *BuildConfig, log *zap.SugaredLogger) error {
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

	return nil
}
