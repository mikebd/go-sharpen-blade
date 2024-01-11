package command

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go-sharpen-blade/observability"
)

const (
	labelCommand = "command"
)

var (
	commandRun = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: observability.AppPrefix + "command_run_total",
		Help: "Command execution count",
	}, []string{labelCommand})
)

func Run(name string) error {
	command, ok := registry[name]
	if !ok {
		return fmt.Errorf("command not found: %s", name)
	}
	commandRun.With(prometheus.Labels{labelCommand: name}).Inc()
	return command()
}
