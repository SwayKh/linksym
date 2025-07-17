package logger

import (
	"fmt"
	"os"

	"github.com/SwayKh/linksym/flags"
)

var noColorIsSet bool

func init() {
	_, noColorIsSet = os.LookupEnv("NO_COLOR")
}

type Attribute string

const (
	SUCCESS Attribute = "\033[1;32m"
	INFO    Attribute = "\033[1;97m"
	WARNING Attribute = "\033[1;33m"
	ERROR   Attribute = "\033[1;31m"
	RESET   Attribute = "\033[0m"
)

func VerboseLog(msgColor Attribute, msg string, args ...any) {
	if noColorIsSet {
		msgColor = RESET
	}
	if *flags.VerboseFlag {
		fmt.Printf(string(msgColor)+"   "+msg+string(RESET)+"\n", args...)
	}
}

func Log(msgColor Attribute, msg string, args ...any) {
	if noColorIsSet {
		msgColor = RESET
	}
	fmt.Printf(string(msgColor)+"   "+msg+string(RESET)+"\n", args...)
}
