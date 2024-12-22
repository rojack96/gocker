package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Top struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (t *Top) DryRun() *Top {
	return &Top{Command: t.Command + option.DryRun()}
}

// ServiceNames - Specify services to remove
func (t *Top) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: t.Command + helpers.ServiceName(serviceNames...)}
}
