package phases

import (
	"github.com/karmek-k/ssgbuild/builder"
)

type Phase interface {
	Perform(b *builder.Build) error
}

func DefaultPhases() []Phase {
	return []Phase{
		ChangeBaseDirPhase{},
		InstallCmdPhase{},
		BuildCmdPhase{},
		CheckResultPhase{},
	}
}