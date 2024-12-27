package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Pause struct {
	command string
}

func NewPause(cmd string) *Pause {
	return &Pause{command: cmd}
}

// DryRun - Execute command in dry run mode
func (p *Pause) DryRun() *Pause {
	return &Pause{command: p.command + option.DryRun()}
}

// ServiceNames - Specify services to list
func (p *Pause) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(p.command + helpers.ServiceName(serviceNames...))
}
