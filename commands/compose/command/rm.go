package command

import (
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	force = "--force"
	stop  = "--stop"
)

type Rm struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (r *Rm) DryRun() *Rm {
	return &Rm{Command: r.Command + option.DryRun()}
}

// Force - Don't ask to confirm removal
func (r *Rm) Force() *Rm {
	return &Rm{Command: r.Command + helpers.Option(force)}
}

// Stop - Stop the containers, if required, before removing
func (r *Rm) Stop() *Rm {
	return &Rm{Command: r.Command + helpers.Option(stop)}
}

// Volumes - Remove any anonymous volumes attached to containers
func (r *Rm) Volumes() *Rm {
	return &Rm{Command: r.Command + option.Volumes()}
}

// ServiceNames - Specify services to remove
func (r *Rm) ServiceNames(serviceNames ...string) *Rm {
	return &Rm{Command: r.Command + helpers.ServiceName(serviceNames...)}
}
