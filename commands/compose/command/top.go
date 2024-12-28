package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
	"github.com/rojack96/gocker/options"
)

type Top struct {
	command string
}

func NewTop(cmd string) *Top {
	return &Top{command: cmd}
}

// DryRun - Execute command in dry run mode
func (t *Top) DryRun() *Top {
	return &Top{command: t.command + options.DryRun()}
}

// ServiceNames - Specify services to remove
func (t *Top) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(t.command + helpers.ServiceName(serviceNames...))
}
