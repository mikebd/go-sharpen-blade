package main

import (
	"flag"
	"go-sharpen-blade/command"
	"go-sharpen-blade/config"
	"go-sharpen-blade/playground/gitPull"
	"log"
	"os"
)

func init() {
	gitPull.Register()
}

func main() {
	defer log.Println("Done")

	args := config.ParseArguments()
	initializeLogging(args.LogTimestamps)
	validateUsage(&args)

	err := run(&args)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(args *config.Arguments) error {
	log.Println("Running:", os.Args)

	if args.WorkingDir != "" {
		currentDir := os.Getenv("PWD")
		err := os.Chdir(args.WorkingDir)
		if err != nil {
			return err
		}
		defer func(dir string) {
			_ = os.Chdir(dir)
		}(currentDir)
	}

	if args.Command != "" {
		err := command.Run(args.Command)
		if err != nil {
			return err
		}
	}

	return nil
}

func initializeLogging(logTimestamps bool) {
	if logTimestamps {
		log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(log.Lshortfile)
	}
	log.SetOutput(os.Stdout)
}

func validateUsage(args *config.Arguments) {
	var invalidUsage bool

	/*
		if len(args.MandatoryStringArgument) == 0 {
			invalidUsage = true
			log.Println("Missing mandatory command line argument: -arg")
		}
	*/

	if invalidUsage {
		flag.Usage()
		os.Exit(2) // Same code used when a flag is provided but not defined
	}
}
