package cmd

import (
	"fmt"
	"os"
)

func Help() {
	fmt.Println("Usage: linksym [subcommand] [flags]")

	fmt.Println("\n Subcommands:")
	fmt.Println("   add [Path] [(optional) Destination]:")
	fmt.Println("     Create a symlink for given path, optionally define a destination for symlink")
	fmt.Println("   remove [Path]")
	fmt.Println("     Remove the symlink and move the file to the original path")

	fmt.Println("\n Flags:")
	fmt.Println("   -h, --help")
	fmt.Println("     Print this help message")
	os.Exit(0)
}
