package config

import (
	"flag"
	"go-sharpen-blade/command"
	"os"
)

type Arguments struct {
	Command       string
	ListCommands  bool
	LogTimestamps bool
	WorkingDir    string
}

func ParseArguments() Arguments {
	args := Arguments{}

	flag.StringVar(&args.Command, "c", "", "Command to run")
	flag.StringVar(&args.WorkingDir, "d", "", "Working directory")
	flag.BoolVar(&args.ListCommands, "l", false, "List commands")
	flag.BoolVar(&args.LogTimestamps, "lt", false, "Log timestamps (UTC)")
	flag.Parse()

	if args.ListCommands {
		command.List()
		os.Exit(0)
	}

	return args
}
