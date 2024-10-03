	"fmt"
	"path/filepath"

// Set values to global variables of InitDirectory and ConfigPath
func SetupDirectories(initDir string, configName string) {
	InitDirectory = ExpandPath(initDir)
	ConfigPath = filepath.Join(InitDirectory, configName)
}

// Set the global HomeDirectory variable. Separated from SetupDirectories to be
// used with the Init Subcommand.
func InitialiseHomePath() error {
	var err error
	HomeDirectory, err = os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Couldn't get the home directory")
	}
	return nil
}
