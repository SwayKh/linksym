package main

import (
	"fmt"
	"os"
)

func main() {
	HOME, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	CONFIG, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	origFilePath := "test/nested/test.tmt"
	if _, err = os.Stat(origFilePath); os.IsNotExist(err) {
		fmt.Println("File Doesn't exist", err)
		os.Exit(1)
	}

	fmt.Println(HOME, CONFIG, cwd, origFilePath)
	err = os.Symlink(cwd+"/"+origFilePath, HOME+"/test.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
