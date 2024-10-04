package flags

import (
	"flag"
)

// Flags are in a different package, since the logger package needs the
// VerboseFlag. Moving the flags logic to commands/linksym.go causes import
// cycle since the subcommands needs logger package and Logger package would
// import commands
var (
	HelpFlag    *bool
	VerboseFlag *bool
)

// Setup the Flags for the CLI
func CreateFlags() {
	// Handle both -h and --help with one boolean
	HelpFlag = flag.Bool("h", false, "Show help")
	flag.BoolVar(HelpFlag, "help", false, "Show help")
	VerboseFlag = flag.Bool("v", false, "Verbose output")
}
