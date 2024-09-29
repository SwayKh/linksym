package cmd

import "github.com/SwayKh/linksym/pkg/logger"

func Help() {
	usage := ` Usage: linksym [flags] [subcommand]

Subcommands:
  init
    Initialize the linksym configuration file (.linksym.yaml) to hold records of symlinks.

  add [path] [destination (Optional) ]
    Create a symlink for the specified path. Optionally define a destination for the symlink.

  remove [path]
    Remove the symlink and restore the original file to its original path.

  source
    Create all symlinks described in the .linksym.yaml configuration file.

  update
    Update the init_directory field and record names in the .linksym.yaml configuration file.

Flags:
  -h, --help
    Display this help message.
  -v
    Show verbose output.`

	logger.Log(usage)
}
