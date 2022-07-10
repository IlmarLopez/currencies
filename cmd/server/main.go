package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/IlmarLopez/currency/internal/config"
	"github.com/IlmarLopez/currency/pkg/log"
)

// Version is the version of the application.
var Version = "1.0.0"

var flagConfig = flag.String("config", "./config/local.yml", "path to the configuration file")

func main() {
	flag.Parse()

	logger, err := log.NewLogger()
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %v", err))
	}

	// load application configurations
	cfg, err := config.Load(*flagConfig, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}

}
