package cmd

import "github.com/SwayKh/linksym/pkg/config"

// Initialise and empty config with cwd as init directory
func Init(configName string) error {
	err := config.InitialiseConfig(configName)
	if err != nil {
		return err
	}
	return nil
}
