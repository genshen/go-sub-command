package version

import (
	"flag"
	"fmt"

	"github.com/genshen/cmds"
)

const VERSION = "0.1.0"

var versionCommand = &cmds.Command{
	Name:        "version",
	Summary:     "print example version",
	Description: "print current example version.",
	CustomFlags: false,
	HasOptions:  false,
}

func init() {
	versionCommand.Runner = &version{}
	fs := flag.NewFlagSet("version", flag.ContinueOnError)
	versionCommand.FlagSet = fs
	versionCommand.FlagSet.Usage = versionCommand.Usage // use default usage provided by cmds.Command.
	cmds.AllCommands = append(cmds.AllCommands, versionCommand)
}

type version struct{}

func (v *version) PreRun() error {
	return nil
}

func (v *version) Run() error {
	fmt.Printf("version\t %s.\n", VERSION)
	fmt.Println("Author\t abc@example.com")
	return nil
}
