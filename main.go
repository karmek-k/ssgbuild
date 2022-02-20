package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/karmek-k/ssgbuild/builder"
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
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	if err := builder.Build(makeCfg(), logger.Sugar()); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
