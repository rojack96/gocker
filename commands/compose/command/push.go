package command

import (
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	ignorePushFailures = "--ignore-push-failures"
)

type Push struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (p *Push) DryRun() *Push {
	return &Push{Command: p.Command + option.DryRun()}
}

// IgnorePushFailures - Push what it can and ignore images with push failures
func (p *Push) IgnorePushFailures() *Push {
	return &Push{Command: p.Command + helpers.Option(ignorePushFailures)}
}

// IncludeDeps - Also push images of services declared as dependencies
func (p *Push) IncludeDeps() *Push {
	return &Push{Command: p.Command + option.IncludeDeps()}
}

// Quiet - Push without printing progress information
func (p *Push) Quiet() *Push {
	return &Push{Command: p.Command + option.Quiet()}
}

// ServiceNames - Specify services to push
func (p *Push) ServiceNames(serviceNames ...string) *Push {
	return &Push{Command: p.Command + helpers.ServiceName(serviceNames...)}
}
