package phases

type PhaseArgs map[string]string

type Phase interface {
	Perform(args PhaseArgs) error
	GetName() string
}

func DefaultPhases() []Phase {
	return []Phase{
		ChangeBaseDirPhase{},
		InstallCmdPhase{},
		BuildCmdPhase{},
		CheckResultPhase{},
	}
}