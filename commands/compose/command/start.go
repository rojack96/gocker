package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Start struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (s *Start) DryRun() *Start {
	return &Start{Command: s.Command + option.DryRun()}
}

// ServiceNames - Specify services to remove
func (s *Start) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: s.Command + helpers.ServiceName(serviceNames...)}
}
