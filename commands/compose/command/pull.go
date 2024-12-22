package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	ignoreBuildable    = "--ignore-buildable"
	ignorePullFailures = "--ignore-pull-failures"
	policy             = "--policy"
)

type Pull struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (p *Pull) DryRun() *Pull {
	return &Pull{Command: p.Command + option.DryRun()}
}

// IgnoreBuildable - Ignore images that can be built
func (p *Pull) IgnoreBuildable() *Pull {
	return &Pull{Command: p.Command + helpers.Option(ignoreBuildable)}
}

// IgnorePullFailures - Pull what it can and ignore images with pull failures
func (p *Pull) IgnorePullFailures() *Pull {
	return &Pull{Command: p.Command + helpers.Option(ignorePullFailures)}
}

// IncludeDeps - Also pull services declared as dependencies
func (p *Pull) IncludeDeps() *Pull {
	return &Pull{Command: p.Command + option.IncludeDeps()}
}

// Policy - Apply pull policy ("missing"|"always")
func (p *Pull) Policy(value string) *Pull {
	return &Pull{Command: p.Command + helpers.String(policy, value)}
}

// Quiet - Pull without printing progress information
func (p *Pull) Quiet() *Pull {
	return &Pull{Command: p.Command + option.Quiet()}
}

// ServiceNames - Specify services to pull
func (p *Pull) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: p.Command + helpers.ServiceName(serviceNames...)}
}
