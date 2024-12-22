package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	noUp = "--no-up"
)

type Watch struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (w *Watch) DryRun() *Watch {
	return &Watch{Command: w.Command + option.DryRun()}
}

// NoUp - Do not build & start services before watching
func (w *Watch) NoUp() *Watch {
	return &Watch{Command: w.Command + helpers.Option(noUp)}
}

func (w *Watch) Quiet() *Watch {
	return &Watch{Command: w.Command + option.Quiet()}
}

// ServiceNames - Specify services to remove
func (w *Watch) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: w.Command + helpers.ServiceName(serviceNames...)}
}
