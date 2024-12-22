package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Restart struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (r *Restart) DryRun() *Restart {
	return &Restart{Command: r.Command + option.DryRun()}
}

// NoDeps - Don't restart dependent services
func (r *Restart) NoDeps() *Restart {
	return &Restart{Command: r.Command + option.NoDeps()}
}

// Timeout - Specify a shutdown timeout in seconds
func (r *Restart) Timeout(seconds int) *Restart {
	return &Restart{Command: r.Command + common.Timeout(seconds)}
}

// ServiceNames - Specify services to push
func (r *Restart) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: r.Command + helpers.ServiceName(serviceNames...)}
}
