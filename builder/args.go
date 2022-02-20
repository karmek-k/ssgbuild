package builder

import (
	"errors"
	"flag"
)

// BuildConfigFromArgs parses CLI arguments and returns
// a pointer to build config or nil, if help should be shown
func BuildConfigFromArgs() (*BuildConfig, error) {
	showHelp := flag.Bool("help", false, "show help")
	baseDir := flag.String("d", "", "the directory to change to for build")
	buildCmd := flag.String("b", "npm run build", "command used for building")
	installCmd := flag.String("i", "npm install", "command used for installing dependencies")
	resultDir := flag.String("r", "dist", "result directory created by the build command")
	flag.Parse()

	// let the outer function handle displaying help
	if *showHelp {
		return nil, nil
	}

	if *baseDir == "" ||
		*buildCmd == "" ||
		*installCmd == "" ||
		*resultDir == "" {
		return nil, errors.New("invalid arguments")
	}

	cfg := BuildConfig{
		BaseDir: *baseDir,
		InstallCmd: *installCmd,
		BuildCmd: *buildCmd,
		ResultDir: *resultDir,
	}

	return &cfg, nil
}