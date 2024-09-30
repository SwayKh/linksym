package main

// Initialise and empty config with cwd as init directory
func Init(configName string) error {
	err := InitialiseConfig(configName)
	if err != nil {
		return err
	}
	return nil
}
