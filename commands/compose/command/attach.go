package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

const (
	detachKeys = "--detach-keys"
	noStdin    = "--no-stdin"
	sigProxy   = "--sig-proxy"
)

type Attach struct {
	command string
}

func NewAttach(cmd string) *Attach {
	return &Attach{command: cmd}
}

// DetachKeys -  Override the key sequence for detaching from a container.
func (a *Attach) DetachKeys(value string) *Attach {
	return &Attach{command: a.command + helpers.String(detachKeys, value)}
}

// DryRun - Execute command in dry run mode
func (a *Attach) DryRun() *Attach {
	return &Attach{command: a.command + common.DryRun()}
}

// Index - index of the container if service has multiple replicas
func (a *Attach) Index(indexOfContainer int) *Attach {
	return &Attach{a.command + common.Index(indexOfContainer)}
}

// NoStdin - Do not attach STDIN
func (a *Attach) NoStdin() *Attach {
	return &Attach{command: a.command + helpers.Option(noStdin)}
}

// SigProxy - Proxy all received signals to the process (default true)
func (a *Attach) SigProxy() *Attach {
	return &Attach{command: a.command + helpers.Option(sigProxy)}
}

func (a *Attach) ServiceName(serviceName string) *common.CommandExecutor {
	return common.SetCommand(a.command + helpers.ServiceName(serviceName))
}
