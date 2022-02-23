package phases

// PhaseArgs is the type for arguments passed to a phase
type PhaseArgs map[string]string

// Phase is a single build phase
// (e.g. installing the dependencies)
type Phase interface {
	Perform(args PhaseArgs) error
	GetName() string
}

// DefaultPhases returns "standard" phases that occur in a usual build
func DefaultPhases() []Phase {
	return []Phase{
		ChangeBaseDirPhase{},
		InstallCmdPhase{},
		BuildCmdPhase{},
		CheckResultPhase{},
	}
}