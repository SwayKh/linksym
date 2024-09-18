package cmd

import "github.com/SwayKh/linksym/pkg/config"

// Initialise and empty config with cwd as init directory
func Init() error {
	err := config.InitialiseConfig()
	if err != nil {
		return err
	}
	return nil
}
