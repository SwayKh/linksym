package main

import (
	"fmt"
)

func VerboseLog(msg string, args ...any) {
	if *VerboseFlag {
		fmt.Printf(msg+"\n", args...)
	}
}

func Log(msg string, args ...any) {
	fmt.Printf(msg+"\n", args...)
}
