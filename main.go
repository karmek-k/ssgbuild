package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/karmek-k/ssgbuild/builder"
)

func main() {
	cfg, err := builder.BuildConfigFromArgs()
	if err != nil {
		os.Exit(1)
	}
	if cfg == nil {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if builder.Build(cfg) != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("site was built successfully")
}
