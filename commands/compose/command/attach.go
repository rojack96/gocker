package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	detachKeys = "--detach-keys"
	noStdin    = "--no-stdin"
	sigProxy   = "--sig-proxy"
)

type Attach struct {
	Command string
}

// DetachKeys -  Override the key sequence for detaching from a container.
func (a *Attach) DetachKeys(value string) *Attach {
	a.Command += helpers.String(detachKeys, value)
	return a
}

// DryRun - Execute command in dry run mode
func (a *Attach) DryRun() *Attach {
	a.Command += option.DryRun()
	return a
}

// Index - index of the container if service has multiple replicas
func (a *Attach) Index(indexOfContainer int) *Attach {
	a.Command += common.Index(indexOfContainer)
	return a
}

// NoStdin - Do not attach STDIN
func (a *Attach) NoStdin() *Attach {
	a.Command += helpers.Option(noStdin)
	return a
}

// SigProxy - Proxy all received signals to the process (default true)
func (a *Attach) SigProxy() *Attach {
	a.Command += helpers.Option(sigProxy)
	return a
}

func (a *Attach) ServiceName(serviceName string) *common.CommandExecutor {
	a.Command += helpers.ServiceName(serviceName)
	return &common.CommandExecutor{Command: a.Command}
}
