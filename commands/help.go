package commands

import "github.com/SwayKh/linksym/logger"

func Help() {
	usage := ` USAGE:
  linksym [flags] [subcommand]

FLAGS:
  -h, --help
    Display this help message.
  -v
    Show verbose output.

AVAILABLE COMMANDS:
  init
    Initialize the linksym configuration file (.linksym.yaml) to hold records of symlinks.

  add [path] [destination (Optional) ]
    Create a symlink for the specified path. Optionally takes a destination path for the symlink.

  remove [path]
    Remove the symlink and restore the original file to its original path.

  source
    Create all symlinks described in the .linksym.yaml configuration file.

  update
    Update the .linksym.yaml configuration file in the current directory.
`

	logger.Log(usage)
}
