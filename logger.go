package main

import (
	"fmt"

	"github.com/SwayKh/linksym/flags"
)

func VerboseLog(msg string, args ...any) {
	if *flags.VerboseFlag {
		fmt.Printf(msg+"\n", args...)
	}
}

func Log(msg string, args ...any) {
	fmt.Printf(msg+"\n", args...)
}
