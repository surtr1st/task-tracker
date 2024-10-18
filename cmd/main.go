package main

import (
	"flag"
	"tasktracker/internal"
)

func main() {
	flag.Parse()
	subcommand := internal.UseSubcommand(flag.Args())
	subcommand.Init()
	subcommand.ParseSubcommands()
}
