package logger

import (
	"github.com/SwayKh/linksym/flags"
	"github.com/fatih/color"
)

const (
	SUCCESS = color.FgGreen
	INFO    = color.FgWhite
	WARNING = color.FgYellow
	ERROR   = color.FgRed
)

func VerboseLog(msgColor color.Attribute, msg string, args ...any) {
	if *flags.VerboseFlag {
		c := color.New(msgColor)
		c.Printf(msg+"\n", args...)
	}
}

func Log(msgColor color.Attribute, msg string, args ...any) {
	c := color.New(msgColor)
	c.Printf(msg+"\n", args...)
}
