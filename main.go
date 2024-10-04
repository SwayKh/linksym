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
		logger.Log("Error: %v", err)
		os.Exit(1)
	}

	App := commands.Application{
		ConfigName:    ".linksym.yaml",
		HomeDirectory: homeDir,
		Configuration: nil, // This is set in Run() function
		ConfigPath:    "",  // This is set in Run() function
		InitDirectory: "",  // This is set in Run() function
	}

	if err := App.Run(); err != nil {
		logger.Log("Error: %v", err)
		os.Exit(1)
	}
}
