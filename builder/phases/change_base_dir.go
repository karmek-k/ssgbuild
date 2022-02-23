package phases

import (
	"errors"
	"fmt"
	"os"

	"github.com/karmek-k/ssgbuild/utils"
)

type ChangeBaseDirPhase struct {}

func (p ChangeBaseDirPhase) Perform(args PhaseArgs) error {
	if err := utils.CheckDir(args["BaseDir"]); err != nil {
		return fmt.Errorf("checking base dir failed: %s", err.Error())
	}

	if os.Chdir(args["BaseDir"]) != nil {
		return errors.New("could not change the base directory")
	}

	return nil
}

func (p ChangeBaseDirPhase) GetName() string {
	return "Change base directory"
}
