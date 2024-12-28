package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

const (
	noUp = "--no-up"
)

type Watch struct {
	command string
}

func NewWatch(cmd string) *Watch {
	return &Watch{command: cmd}
}

// DryRun - Execute command in dry run mode
func (w *Watch) DryRun() *Watch {
	return &Watch{command: w.command + common.DryRun()}
}

// NoUp - Do not build & start services before watching
func (w *Watch) NoUp() *Watch {
	return &Watch{command: w.command + helpers.Option(noUp)}
}

func (w *Watch) Quiet() *Watch {
	return &Watch{command: w.command + common.Quiet()}
}

// ServiceNames - Specify services to remove
func (w *Watch) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(w.command + helpers.ServiceName(serviceNames...))

}
