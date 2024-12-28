package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

const (
	orphans = "--orphans"
	status  = "--status"
)

type Ps struct {
	command string
}

func NewPs(cmd string) *Ps {
	return &Ps{command: cmd}
}

// All - Show all stopped containers (including those created by the run command)
func (p *Ps) All() *Ps {
	return &Ps{command: p.command + common.All()}
}

// DryRun - Execute command in dry run mode
func (p *Ps) DryRun() *Ps {
	return &Ps{command: p.command + common.DryRun()}
}

// Filter - Filter services by a property (supported filters: status)
func (p *Ps) Filter(condition string) *Ps {
	return &Ps{command: p.command + common.Filter(condition)}
}

// Format - Format output using a custom template
func (p *Ps) Format(formatValue string) *Ps {
	return &Ps{command: p.command + common.Format(formatValue)}
}

// NoTrunc - Don't truncate output
func (p *Ps) NoTrunc() *Ps {
	return &Ps{command: p.command + common.NoTrunc()}
}

// Orphans - Include orphaned services (not declared by project) (default true)
func (p *Ps) Orphans() *Ps {
	return &Ps{command: p.command + helpers.Option(orphans)}
}

// Services - Display services
func (p *Ps) Services() *Ps {
	return &Ps{command: p.command + common.ServicesOption()}
}

// Status - Filter services by status (e.g., paused, running)
func (p *Ps) Status(statusValues ...string) *Ps {
	return &Ps{command: p.command + helpers.StringArray(status, statusValues...)}
}

// Quiet - Only display IDs
func (p *Ps) Quiet() *Ps {
	return &Ps{command: p.command + common.Quiet()}
}

// ServiceNames - Specify services to list
func (p *Ps) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(p.command + helpers.ServiceName(serviceNames...))
}
