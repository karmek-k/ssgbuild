package phases

import (
	"fmt"

	"github.com/karmek-k/ssgbuild/pkg/utils"
)

type InstallCmdPhase struct{}

func (p InstallCmdPhase) Perform(args PhaseArgs) error {
	installOut, err := utils.StringToCmd(args["InstallCmd"]).CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"failed running the install command\n%s",
			string(installOut),
		)
	}

	return nil
}

func (p InstallCmdPhase) GetName() string {
	return "Execute the install command"
}
