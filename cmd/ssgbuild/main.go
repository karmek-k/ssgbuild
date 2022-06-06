package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/karmek-k/ssgbuild/pkg/builder"
	"github.com/karmek-k/ssgbuild/pkg/builder/phases"
	"go.uber.org/zap"
)

func makeCfg() *builder.BuildConfig {
	cfg, err := builder.BuildConfigFromArgs()
	if err != nil {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if cfg == nil {
		flag.PrintDefaults()
		os.Exit(0)
	}

	return cfg
}

func main() {
	unsugaredLog, _ := zap.NewDevelopment()
	defer unsugaredLog.Sync()

	cfg := makeCfg()
	log := unsugaredLog.Sugar()

	build := builder.Build{
		Cfg: cfg,
		Log: log,
	}

	if err := build.Start(phases.DefaultPhases()); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
