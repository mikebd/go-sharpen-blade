package config

import (
	"flag"
	"go-sharpen-blade/internal/command"
	"os"
)

type Arguments struct {
	Command        string
	ListCommands   bool
	LogTimestamps  bool
	O11yPrometheus bool
	WorkingDir     string
}

func ParseArguments() Arguments {
	args := Arguments{}

	flag.StringVar(&args.Command, "c", "", "Command to run")
	flag.StringVar(&args.WorkingDir, "d", "", "Working directory")
	flag.BoolVar(&args.ListCommands, "l", false, "List commands")
	flag.BoolVar(&args.LogTimestamps, "lt", false, "Log timestamps (UTC)")
	flag.BoolVar(&args.O11yPrometheus, "o11y-prometheus", false, "Enable Prometheus metrics")
	flag.Parse()

	if args.ListCommands {
		command.List()
		os.Exit(0)
	}

	return args
}
