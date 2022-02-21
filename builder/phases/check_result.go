package phases

import (
	"fmt"

	"github.com/karmek-k/ssgbuild/builder"
	"github.com/karmek-k/ssgbuild/utils"
)

type CheckResultPhase struct {}

func (p CheckResultPhase) Perform(b *builder.Build) error {
	b.Log.Debugw("checking the result dir",
		"dir", b.Cfg.ResultDir,
	)

	if err := utils.CheckDir(b.Cfg.ResultDir); err != nil {
		b.Log.Errorw("checking result dir failed",
			"name", b.Cfg.Name,
			"dir", b.Cfg.ResultDir,
		)

		return fmt.Errorf("checking result dir failed: %s", err.Error())
	}

	return nil
}