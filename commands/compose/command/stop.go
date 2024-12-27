package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Stop struct {
	command string
}

func NewStop(cmd string) *Stop {
	return &Stop{command: cmd}
}

// DryRun - Execute command in dry run mode
func (s *Stop) DryRun() *Stop {
	return &Stop{command: s.command + option.DryRun()}
}

// Timeout - Specify a shutdown timeout in seconds
func (s *Stop) Timeout(seconds int) *Stop {
	return &Stop{command: s.command + common.Timeout(seconds)}
}

// ServiceNames - Specify services to remove
func (s *Stop) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(s.command + helpers.ServiceName(serviceNames...))
}
