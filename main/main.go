package main

import (
	"flag"
	"github.com/mikebd/go-util/pkg/directory"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-sharpen-blade/command"
	"go-sharpen-blade/config"
	"go-sharpen-blade/playground/aws"
	"go-sharpen-blade/playground/gitPull"
	sudoku "go-sharpen-blade/playground/sudoku/game"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	prometheusPath = "/metrics"
	prometheusPort = ":2112"
)

func init() {
	aws.Register()
	gitPull.Register()
	sudoku.Register()
}

func main() {
	start := time.Now()
	defer func() {
		log.Println("Done in", time.Since(start).Round(time.Millisecond))
	}()

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
		_, restoreDir, err := directory.ChangeDirectory(args.WorkingDir)
		if err != nil {
			return err
		}
		defer restoreDir()
	}

	if args.Command != "" {
		if args.O11yPrometheus {
			go func() {
				http.Handle(prometheusPath, promhttp.Handler())
				_ = http.ListenAndServe(prometheusPort, nil)
			}()
		}
		err := command.Run(args.Command, args.O11yPrometheus)
		if err != nil {
			return err
		}

		if args.O11yPrometheus {
			// Consider making this delay configurable
			prometheusDelaySeconds := 20
			log.Printf("Waiting %d seconds for Prometheus to scrape before exiting...\n", prometheusDelaySeconds)
			time.Sleep(time.Duration(prometheusDelaySeconds) * time.Second)
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
