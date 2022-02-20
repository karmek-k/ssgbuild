package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("ssgbuild")

	showHelp := flag.Bool("help", false, "show help")
	baseDir := flag.String("d", "", "the directory to change to for build")
	command := flag.String("c", "", "command used for building")
	resultDir := flag.String("r", "", "result directory created by the build command")
	flag.Parse()

	if *showHelp || *baseDir == "" || *command == "" || *resultDir == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Printf(
		"baseDir=%s command=%s resultDir=%s\n",
		*baseDir,
		*command,
		*resultDir,
	)
}
