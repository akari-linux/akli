package main

import (
	"github.com/alecthomas/kong"
)

const AkliVersion string = "0.0.1pa"

var CLI struct {
	Verbose bool             `short:"v" help:"Enable verbosity."`
	Version kong.VersionFlag `short:"V" help:"Print version and exit."`
}

func main() {
	ctx := kong.Parse(&CLI, kong.Vars{
		"version": AkliVersion,
	})
	switch ctx.Command() {
	default:
		ctx.PrintUsage(false)
	}
}
