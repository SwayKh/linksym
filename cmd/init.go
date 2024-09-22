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

// Update the Init Directory variable in the config
func UpdateInit(configuration *config.AppConfig, configName string) error {
	err := config.UpdateInitDirectory(configuration, configName)
	if err != nil {
		return err
	}
	return nil
}
