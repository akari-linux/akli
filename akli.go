package main

import (
	"github.com/alecthomas/kong"
)

var CLI struct {
	Verbose bool `short:"v" help:"Enable verbosity."`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	default:
		ctx.PrintUsage(false)
	}
}
