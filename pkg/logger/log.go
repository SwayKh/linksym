package logger

import (
	"fmt"

	"github.com/SwayKh/linksym/pkg/flags"
)

func Log(msg string, args ...any) {
	if *flags.VerboseFlag {
		fmt.Printf(msg, args...)
	}
}

func VerboseLog(msg string, args ...any) {
	fmt.Printf(msg, args...)
}
