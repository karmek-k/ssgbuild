package phases

type Phase interface {
	Perform(args map[string]string) error
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