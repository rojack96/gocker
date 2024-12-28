package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
	"github.com/rojack96/gocker/options"
)

const (
	ignorePushFailures = "--ignore-push-failures"
)

type Push struct {
	command string
}

func NewPush(cmd string) *Push {
	return &Push{command: cmd}
}

// DryRun - Execute command in dry run mode
func (p *Push) DryRun() *Push {
	return &Push{command: p.command + options.DryRun()}
}

// IgnorePushFailures - Push what it can and ignore images with push failures
func (p *Push) IgnorePushFailures() *Push {
	return &Push{command: p.command + helpers.Option(ignorePushFailures)}
}

// IncludeDeps - Also push images of services declared as dependencies
func (p *Push) IncludeDeps() *Push {
	return &Push{command: p.command + options.IncludeDeps()}
}

// Quiet - Push without printing progress information
func (p *Push) Quiet() *Push {
	return &Push{command: p.command + options.Quiet()}
}

// ServiceNames - Specify services to push
func (p *Push) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(p.command + helpers.ServiceName(serviceNames...))
}
