package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

type Unpause struct {
	command string
}

func NewUnpause(cmd string) *Unpause {
	return &Unpause{command: cmd}
}

// DryRun - Execute command in dry run mode
func (u *Unpause) DryRun() *Unpause {
	return &Unpause{command: u.command + common.DryRun()}
}

// ServiceNames - Specify services to remove
func (u *Unpause) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(u.command + helpers.ServiceName(serviceNames...))
}
