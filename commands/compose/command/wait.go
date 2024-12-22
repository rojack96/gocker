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
	Command string
}

// DownProject - Drops project when the first container stops
func (w *Wait) DownProject() *Wait {
	return &Wait{Command: w.Command + helpers.Option(downProject)}
}

// DryRun - Execute command in dry run mode
func (w *Wait) DryRun() *Wait {
	return &Wait{Command: w.Command + option.DryRun()}
}

// ServiceNames - Specify services to remove
func (w *Wait) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: w.Command + helpers.ServiceName(serviceNames...)}
}
