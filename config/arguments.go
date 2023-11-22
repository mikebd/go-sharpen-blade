package config

import (
	"flag"
	"fmt"
	"go-sharpen-blade/command"
	"os"
)

type Arguments struct {
	Command       string
	ListCommands  bool
	LogTimestamps bool
}

func ParseArguments() Arguments {
	args := Arguments{}

	flag.StringVar(&args.Command, "c", "", "Command to run")
	flag.BoolVar(&args.ListCommands, "l", false, "List commands")
	flag.BoolVar(&args.LogTimestamps, "lt", false, "Log timestamps (UTC)")
	flag.Parse()

	if args.ListCommands {
		command.List()
		os.Exit(0)
	}

	if args.Command != "" {
		err := command.Run(args.Command)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	return args
}
