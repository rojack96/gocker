package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
	"github.com/rojack96/gocker/options"
)

type Start struct {
	command string
}

func NewStart(cmd string) *Start {
	return &Start{command: cmd}
}

// DryRun - Execute command in dry run mode
func (s *Start) DryRun() *Start {
	return &Start{command: s.command + options.DryRun()}
}

// ServiceNames - Specify services to remove
func (s *Start) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(s.command + helpers.ServiceName(serviceNames...))
}
