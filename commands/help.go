package commands

import (
	"fmt"
)

const (
	white              = "\033[97m"
	boldWhite          = "\033[1;97m"
	boldUnderlineWhite = "\033[1;4;97m"
	reset              = "\033[0m"
)

func print(msg, style string, tabs int) {
	switch tabs {
	case 0:
		fmt.Println(style + msg + reset)
	case 1:
		fmt.Println("  " + style + msg + reset)
	case 2:
		fmt.Println("    " + style + msg + reset)
	default:
		fmt.Println("Undefined number of tabs supplied in help message")
	}
}

func Help() {
	print("USAGE:", boldUnderlineWhite, 0)
	print("linksym [flags] [subcommand]", boldWhite, 1)
	print("", white, 0)
	print("FLAGS:", boldUnderlineWhite, 0)
	print("-h, --help", boldWhite, 1)
	print("Display this help message.", white, 2)
	print("-v", boldWhite, 1)
	print("Show verbose output.", white, 2)
	print("", white, 0)
	print("AVAILABLE COMMANDS:", boldUnderlineWhite, 0)
	print("init", boldWhite, 1)
	print("Initialize the linksym configuration file (.linksym.yaml) to hold records of symlinks.", white, 2)
	print("", white, 0)
	print("add [target] [destination (Optional)]", boldWhite, 1)
	print("Create a symlink for the specified path. Optionally takes a destination path for the symlink.", white, 2)
	print("", white, 0)
	print("record [target] [destination (Optional)]", boldWhite, 1)
	print("Creates a record of symlink in .linksym.yaml, without actually creating symlink.", white, 2)
	print("", white, 0)
	print("remove [target(s)...]", boldWhite, 1)
	print("Remove the symlink and restore the original file to its original path.", white, 2)
	print("", white, 0)
	print("restore [target(s)...]", boldWhite, 1)
	print("Create symlink for specified target(s) that has a record in .linksym.yaml configuration file.", white, 2)
	print("", white, 0)
	print("source", boldWhite, 1)
	print("Create all symlinks described in the .linksym.yaml configuration file.", white, 2)
	print("", white, 0)
	print("update", boldWhite, 1)
	print("Update the .linksym.yaml configuration file in the current directory.", white, 2)
	print("", white, 0)
}
