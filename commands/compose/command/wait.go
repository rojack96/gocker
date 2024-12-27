package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	downProject = "--down-project"
)

type Wait struct {
	command string
}

func NewWait(cmd string) *Wait {
	return &Wait{command: cmd}
}

// DownProject - Drops project when the first container stops
func (w *Wait) DownProject() *Wait {
	return &Wait{command: w.command + helpers.Option(downProject)}
}

// DryRun - Execute command in dry run mode
func (w *Wait) DryRun() *Wait {
	return &Wait{command: w.command + option.DryRun()}
}

// ServiceNames - Specify services to remove
func (w *Wait) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: w.command + helpers.ServiceName(serviceNames...)}
}
