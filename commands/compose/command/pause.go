package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
	"github.com/rojack96/gocker/options"
)

type Pause struct {
	command string
}

func NewPause(cmd string) *Pause {
	return &Pause{command: cmd}
}

// DryRun - Execute command in dry run mode
func (p *Pause) DryRun() *Pause {
	return &Pause{command: p.command + options.DryRun()}
}

// ServiceNames - Specify services to list
func (p *Pause) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(p.command + helpers.ServiceName(serviceNames...))
}
