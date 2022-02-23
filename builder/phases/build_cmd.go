package phases

import (
	"fmt"

	"github.com/karmek-k/ssgbuild/utils"
)

type BuildCmdPhase struct {}

func (p BuildCmdPhase) Perform(args map[string]string) error {
	buildOut, err := utils.StringToCmd(args["BuildCmd"]).CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"failed running the build command\n%s",
			string(buildOut),
		)
	}

	return nil
}

func (p BuildCmdPhase) GetName() string {
	return "Execute the build command"
}