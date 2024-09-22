package cmd

import "fmt"

func Help() {
	usage := ` Usage: linksym [subcommand] [flags]

Subcommands:
  init [flags]
    Initialize the linksym configuration file (.linksym.yaml) to hold records of symlinks.
    Flags:
      -u, --update
        Update the init_directory field in the .linksym.yaml configuration file.

  add [path] [(optional) destination]
    Create a symlink for the specified path. Optionally define a destination for the symlink.

  remove [path]
    Remove the symlink and restore the original file to its original path.

  source
    Create all symlinks described in the .linksym.yaml configuration file.

Flags:
  -h, --help
    Display this help message.`

	fmt.Println(usage)
}
