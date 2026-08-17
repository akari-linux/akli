package main

import (
	"os"
	"runtime"

	"charm.land/log/v2"
	"github.com/alecthomas/kong"
)

const AkliVersion string = "0.0.2pa"

var CLI struct {
	Verbose bool             `short:"v" help:"Enable verbosity."`
	Version kong.VersionFlag `short:"V" help:"Print version and exit."`

	Drivers struct {
		Install struct {
		} `cmd:"" help:"Automatically installs any required graphics drivers"`

		Remove struct {
		} `cmd:"" help:"Automatically removes installed graphics drivers"`
	} `cmd:"" help:"Manage your system's drivers."`
}

func main() {
	var logger = log.New(os.Stderr)
	logger.SetLevel(log.DebugLevel)

	ctx := kong.Parse(&CLI, kong.Vars{
		"version": AkliVersion,
	})

	if CLI.Verbose {
		logger.Debugf("Running akli on %s with %s", runtime.GOOS, runtime.Version())
	}

	switch ctx.Command() {
	case "drivers install":
		// TODO: Implement automatic GPU driver installation
	case "drivers remove":
		// TODO: Implement automatic GPU driver removal
	}
}
