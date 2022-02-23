package phases

import (
	"fmt"

	"github.com/karmek-k/ssgbuild/utils"
)

type CheckResultPhase struct {}

func (p CheckResultPhase) Perform(args PhaseArgs) error {
	if err := utils.CheckDir(args["ResultDir"]); err != nil {
		return fmt.Errorf("checking result dir failed: %s", err.Error())
	}

	return nil
}

func (p CheckResultPhase) GetName() string {
	return "Check the result directory"
}