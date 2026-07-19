package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
	"github.com/rojack96/gocker/options"
)

const (
	force = "--force"
	stop  = "--stop"
)

type Rm struct {
	command string
}

func NewRm(cmd string) *Rm {
	return &Rm{command: cmd}
}

// DryRun - Execute command in dry run mode
func (r *Rm) DryRun() *Rm {
	return &Rm{command: r.command + options.DryRun()}
}

// Force - Don't ask to confirm removal
func (r *Rm) Force() *Rm {
	return &Rm{command: r.command + helpers.Option(force)}
}

// Stop - Stop the containers, if required, before removing
func (r *Rm) Stop() *Rm {
	return &Rm{command: r.command + helpers.Option(stop)}
}

// Volumes - Remove any anonymous volumes attached to containers
func (r *Rm) Volumes() *Rm {
	return &Rm{command: r.command + options.Volumes()}
}

// ServiceNames - Specify services to remove
func (r *Rm) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(r.command + helpers.ServiceName(serviceNames...))
}
