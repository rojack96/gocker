package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Unpause struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (u *Unpause) DryRun() *Unpause {
	return &Unpause{Command: u.Command + option.DryRun()}
}

// ServiceNames - Specify services to remove
func (u *Unpause) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: u.Command + helpers.ServiceName(serviceNames...)}
}
