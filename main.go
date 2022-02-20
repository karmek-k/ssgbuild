package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/karmek-k/ssgbuild/builder"
)

func makeCfg() *builder.BuildConfig {
	cfg, err := builder.BuildConfigFromArgs()
	if err != nil {
		os.Exit(1)
	}
	if cfg == nil {
		flag.PrintDefaults()
		os.Exit(0)
	}

	return cfg
}

func main() {
	if err := builder.Build(makeCfg()); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("site was built successfully")
}
