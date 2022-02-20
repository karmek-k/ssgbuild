package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/karmek-k/ssgbuild/builder"
)

func main() {
	showHelp := flag.Bool("help", false, "show help")
	baseDir := flag.String("d", "", "the directory to change to for build")
	buildCmd := flag.String("b", "npm run build", "command used for building")
	installCmd := flag.String("i", "npm install", "command used for installing dependencies")
	resultDir := flag.String("r", "dist", "result directory created by the build command")
	flag.Parse()

	if *showHelp ||
		*baseDir == "" ||
		*buildCmd == "" ||
		*installCmd == "" ||
		*resultDir == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	cfg := builder.BuildConfig{
		BaseDir: *baseDir,
		InstallCmd: *installCmd,
		BuildCmd: *buildCmd,
		ResultDir: *resultDir,
	}

	err := builder.Build(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("site was built successfully")
}
