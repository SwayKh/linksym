# Linksym - A Symlink-ing tool

`linksym` is a dotfiles management tool which acts as a wrapper around `ln` for
creating symbolic links (symlinks) and creating a record of those symlinks in a
configuration file, allowing to easily recreate those symlinks later from a
single file using a simple command.

## Features

- Allows easy creation of symlinks.
- Automatically records symlinks in `.linksym.yaml` file
- Allow you to recreate or restore symlinks from the configuration file

## Installation

Make sure you have `go` installed on your system.

Install it directly using `go`

```bash
go install github.com/SwayKh/linksym@latest
```

Uninstall `linksym` by running:

```bash
$ rm -f $(which linksym)
```

## Usage

```
linksym init
```

Creates a `.linksym.yaml` file in the current directory This file acts as a
database for storing record of symlinks. All other commands require for the
`.linksym.yaml` file to be present and hence this command is required to be
run before any other command.

```
linksym add [target] [destination (optional)]
```

Moves the file from `target-path` to `destination-path` (Or the current
directory if no destination path is provided) and creates a symlinks at source
pointing to destination. And records it in `.linksym.yaml`.

> [!NOTE]
> the `linksym add` command can also be used in a way similar to `ln` where if
> the target directory or file is already moved to the destination, running
> `linksym add [symlink location] [target path]` will create a symlink there
> anyway.

```
linksym record [target] [destination (optional)]
```

Separate command to add a symlink record to `.linksym.yaml` file. Skips the
Moving and symlinking step of the `add` subcommand. Useful for creating a record
of symlink paths that are already present on the system.

```
linksym remove [target(s)...]
```

Removes the symlink and restores the target file or directory to its original
path and remove the record from `.linksym.yaml`.

```
linksym update
```

Updates the `.linksym.yaml` file in the current directory. This updates the Init
directory field in `.linksym.yaml` file with the current directory. and updates the
`record name` fields appropriately.

```
linksym source
```

Reads the `.linksym.yaml` file in the current directory and creates symlinks for
each record. Useful for replicating recorded symlinks on a different
system or machine.

> [!WARNING]
> Using the source command to create symlinks will overwrite any existing
> Directory or File at the Source path where the symlink will be made

#### Help

```
$ linksym -h/ --help

USAGE:
  linksym [flags] [subcommand]

FLAGS:
  -h, --help
    Display this help message.
  -v
    Show verbose output.

AVAILABLE COMMANDS:
  init
    Initialize the linksym configuration file (.linksym.yaml) to hold records of symlinks.

  add [target] [destination (Optional)]
    Create a symlink for the specified path. Optionally takes a destination path for the symlink.

  record [target] [destination (Optional)]
    Creates a record of symlink in .linksym.yaml, which actually creating symlink.

  remove [target(s)...]
    Remove the symlink and restore the original file to its original path.

  source
    Create all symlinks described in the .linksym.yaml configuration file.

  update
    Update the .linksym.yaml configuration file in the current directory.
```

## Motivation

I know that there are quite a few tools out there for managing dotfiles. Like
all these utilities on [http://dotfiles.github.io/utilities/]. Stow and
Chezmoi are famous choices for dotfiles management. But they have completely
different workflows and I never got used to stow with the packages of dotfiles.

I manage my dotfiles with a simple bash script which just has `ln` command to
link each of my [dotfiles](https://github.com/swaykh/dotfiles). This project was
made to ease that process and make it easier to manage symlinks.

#### Workflow

> - Make a dotfiles directory.
> - Initialize linksym with `linksym init`.
> - Add dotfiles to dotfiles directory with `linksym add`.
> - Remove any unneeded dotfiles with `linksym remove`.
> - Track changes with git.
> - On another system or machine, Clone your dotfiles repo.
> - Run `linksym update` to update the `.linksym.yaml` file.
> - Run `linksym source` to create symlinks from the `.linksym.yaml` file.
> - Profit.

## License

[The MIT License ](./LICENSE)
