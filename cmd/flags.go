package cmd

import (
	"flag"
)

var (
	AddFlag        *flag.FlagSet
	RemoveFlag     *flag.FlagSet
	InitFlag       *flag.FlagSet
	SourceFlag     *flag.FlagSet
	UpdateInitBool bool
	HelpFlag       *bool
	SPath          string
	DPath          string
	RemovePath     string
	ConfigPath     string
)

// Setup the Flags for the CLI
func CreateFlags() {
	// Handle both -h and --help with one boolean
	HelpFlag = flag.Bool("h", false, "Show help")
	flag.BoolVar(HelpFlag, "help", false, "Show help")

	AddFlag = flag.NewFlagSet("add", flag.ExitOnError)
	RemoveFlag = flag.NewFlagSet("remove", flag.ExitOnError)
	InitFlag = flag.NewFlagSet("init", flag.ExitOnError)
	SourceFlag = flag.NewFlagSet("source", flag.ExitOnError)

	AddFlag.StringVar(&SPath, "source", "", "Source path for the file to symlink")
	AddFlag.StringVar(&DPath, "destination", "", "(Optional) Destination for symlink")

	RemoveFlag.StringVar(&RemovePath, "path", "", "Path to remove symlink")

	InitFlag.BoolVar(&UpdateInitBool, "u", false, "Update just the Init Directory variable on config")
	InitFlag.BoolVar(&UpdateInitBool, "update", false, "Update just the Init Directory variable on config")

	SourceFlag.StringVar(&ConfigPath, "path", "", "Path to config file to source")
}
