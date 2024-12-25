package commands

import (
	"github.com/fatih/color"
)

func Help() {
	white := color.New(color.FgWhite).PrintlnFunc()
	boldWhite := color.New(color.FgWhite, color.Bold).PrintlnFunc()
	underlineBoldWhite := color.New(color.FgWhite, color.Bold, color.Underline).PrintlnFunc()

	underlineBoldWhite("USAGE:")
	boldWhite("  linksym [flags] [subcommand]")
	white()
	underlineBoldWhite("FLAGS:")
	boldWhite("  -h, --help")
	white("    Display this help message.")
	boldWhite("  -v")
	white("    Show verbose output.")
	white()
	underlineBoldWhite("AVAILABLE COMMANDS:")
	boldWhite("  init")
	white("    Initialize the linksym configuration file (.linksym.yaml) to hold records of symlinks.")
	white()
	boldWhite("  add [target] [destination (Optional)]")
	white("    Create a symlink for the specified path. Optionally takes a destination path for the symlink.")
	white()
	boldWhite("  record [target] [destination (Optional)]")
	white("    Creates a record of symlink in .linksym.yaml, without actually creating symlink.")
	white()
	boldWhite("  remove [target(s)...]")
	white("    Remove the symlink and restore the original file to its original path.")
	white()
	boldWhite("  restore [target(s)...]")
	white("    Create symlink for specified target(s) that has a record in .linksym.yaml configuration file.")
	white()
	boldWhite("  source")
	white("    Create all symlinks described in the .linksym.yaml configuration file.")
	white()
	boldWhite("  update")
	white("    Update the .linksym.yaml configuration file in the current directory.")
	white()
}
