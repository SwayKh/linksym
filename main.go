package main

import (
	"fmt"
	"os"

	"github.com/SwayKh/linksym/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
