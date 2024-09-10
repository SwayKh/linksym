package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/SwayKh/linksym/pkg/config"
)

var (
	AddFlag    *flag.FlagSet
	RemoveFlag *flag.FlagSet
	InitFlag   *flag.FlagSet
	HelpFlag   *bool
	SPath      string
	DPath      string
	RemovePath string
)

func CreateFlags() {
	// Handle both -h and --help with one boolean
	HelpFlag = flag.Bool("h", false, "Show help")
	flag.BoolVar(HelpFlag, "help", false, "Show help")

	AddFlag = flag.NewFlagSet("add", flag.ExitOnError)
	RemoveFlag = flag.NewFlagSet("remove", flag.ExitOnError)
	InitFlag = flag.NewFlagSet("init", flag.ExitOnError)

	AddFlag.StringVar(&SPath, "source", "", "Source path for the file to symlink")
	AddFlag.StringVar(&DPath, "destination", "", "(Optional) Destination for symlink")

	RemoveFlag.StringVar(&RemovePath, "path", "", "Path to remove symlink")
}

func Init() error {
	err := config.InitialiseConfig()
	if err != nil {
		fmt.Printf("Error initialising config: %s\n", err)
		os.Exit(1)
	}
	return nil
}

func Add() error {
	return nil
}

func Remove() error {
	return nil
}

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
	os.Exit(1)
}
