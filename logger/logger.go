package logger

import (
	"fmt"

	"github.com/SwayKh/linksym/flags"
)

type Attribute string

const (
	SUCCESS            Attribute = "\033[1;32m"
	INFO               Attribute = "\033[1;97m"
	WARNING            Attribute = "\033[1;33m"
	ERROR              Attribute = "\033[1;31m"
	BOLDUNDERLINEWHITE Attribute = "\003[1;4;37m"
)

func VerboseLog(msgColor Attribute, msg string, args ...any) {
	if *flags.VerboseFlag {
		fmt.Printf(string(msgColor)+"   "+msg+"\n", args...)
	}
}

func Log(msgColor Attribute, msg string, args ...any) {
	fmt.Printf(string(msgColor)+"   "+msg+"\n", args...)
}
