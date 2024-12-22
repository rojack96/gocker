package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Pause struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (p *Pause) DryRun() *Pause {
	return &Pause{Command: p.Command + option.DryRun()}
}

// ServiceNames - Specify services to list
func (p *Pause) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: p.Command + helpers.ServiceName(serviceNames...)}
}
