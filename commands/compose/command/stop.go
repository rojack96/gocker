package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Stop struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (s *Stop) DryRun() *Stop {
	return &Stop{Command: s.Command + option.DryRun()}
}

// Timeout - Specify a shutdown timeout in seconds
func (s *Stop) Timeout(seconds int) *Stop {
	return &Stop{Command: s.Command + common.Timeout(seconds)}
}

// ServiceNames - Specify services to remove
func (s *Stop) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: s.Command + helpers.ServiceName(serviceNames...)}
}
