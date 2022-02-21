package phases

import (
	"errors"
	"fmt"
	"os"

	"github.com/karmek-k/ssgbuild/builder"
	"github.com/karmek-k/ssgbuild/utils"
)

type ChangeBaseDirPhase struct {}

func (p ChangeBaseDirPhase) Perform(b *builder.Build) error {
	b.Log.Debugw("changing to the base directory",
		"dir", b.Cfg.BaseDir,
	)
	
	if err := utils.CheckDir(b.Cfg.BaseDir); err != nil {
		b.Log.Errorw("checking base dir failed",
			"name", b.Cfg.Name,
			"dir", b.Cfg.BaseDir,
			"error", err.Error(),
		)

		return fmt.Errorf("checking base dir failed: %s", err.Error())
	}

	if os.Chdir(b.Cfg.BaseDir) != nil {
		b.Log.Errorw("changing to base dir failed",
			"name", b.Cfg.Name,
			"dir", b.Cfg.BaseDir,
		)

		return errors.New("could not change the base directory")
	}

	return nil
}
