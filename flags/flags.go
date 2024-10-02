package flags

import (
	"flag"
)

var (
	AddFlag     *flag.FlagSet
	RemoveFlag  *flag.FlagSet
	InitFlag    *flag.FlagSet
	SourceFlag  *flag.FlagSet
	UpdateFlag  *flag.FlagSet
	HelpFlag    *bool
	VerboseFlag *bool
)

// Setup the Flags for the CLI
func CreateFlags() {
	// Handle both -h and --help with one boolean
	HelpFlag = flag.Bool("h", false, "Show help")
	flag.BoolVar(HelpFlag, "help", false, "Show help")

	VerboseFlag = flag.Bool("v", false, "Verbose output")

	AddFlag = flag.NewFlagSet("add", flag.ExitOnError)
	RemoveFlag = flag.NewFlagSet("remove", flag.ExitOnError)
	InitFlag = flag.NewFlagSet("init", flag.ExitOnError)
	SourceFlag = flag.NewFlagSet("source", flag.ExitOnError)
	UpdateFlag = flag.NewFlagSet("update", flag.ExitOnError)
}
