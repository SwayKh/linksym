package main

import (
	"os"

	"github.com/SwayKh/linksym/commands"
	"github.com/SwayKh/linksym/config"
	"github.com/SwayKh/linksym/logger"
)

func main() {
	homeDir, err := config.InitialiseHomePath()
	if err != nil {
		logger.Log(logger.ERROR, "Error: %v", err)
		os.Exit(1)
	}

	App := commands.Application{
		ConfigName:    ".linksym.yaml",
		HomeDirectory: homeDir,
		Configuration: nil, // This is set in Run() function in linksym.go
		ConfigPath:    "",  // This is set in Run() function in linksym.go
		InitDirectory: "",  // This is set in Run() function in linksym.go
	}

	if err := App.Run(); err != nil {
		logger.Log(logger.ERROR, "Error: %v", err)
		os.Exit(1)
	}
}
