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
	buildCmd := flag.String("c", "", "command used for building")
	installCmd := flag.String("i", "npm install", "command used for installing dependencies")
	resultDir := flag.String("r", "", "result directory created by the build command")
	flag.Parse()

	if *showHelp ||
		*baseDir == "" ||
		*buildCmd == "" ||
		*installCmd == "" ||
		*resultDir == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	err := builder.Build(*baseDir, *installCmd, *buildCmd, *resultDir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("site was built successfully")
}
