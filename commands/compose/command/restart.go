package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
	"github.com/rojack96/gocker/options"
)

type Restart struct {
	command string
}

func NewRestart(cmd string) *Restart {
	return &Restart{command: cmd}
}

// DryRun - Execute command in dry run mode
func (r *Restart) DryRun() *Restart {
	return &Restart{command: r.command + options.DryRun()}
}

// NoDeps - Don't restart dependent services
func (r *Restart) NoDeps() *Restart {
	return &Restart{command: r.command + options.NoDeps()}
}

// Timeout - Specify a shutdown timeout in seconds
func (r *Restart) Timeout(seconds int) *Restart {
	return &Restart{command: r.command + common.Timeout(seconds)}
}

// ServiceNames - Specify services to push
func (r *Restart) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(r.command + helpers.ServiceName(serviceNames...))
}
