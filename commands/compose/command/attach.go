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
	return &Attach{Command: a.Command + helpers.String(detachKeys, value)}
}

// DryRun - Execute command in dry run mode
func (a *Attach) DryRun() *Attach {
	return &Attach{Command: a.Command + option.DryRun()}
}

// Index - index of the container if service has multiple replicas
func (a *Attach) Index(indexOfContainer int) *Attach {
	return &Attach{a.Command + common.Index(indexOfContainer)}
}

// NoStdin - Do not attach STDIN
func (a *Attach) NoStdin() *Attach {
	return &Attach{Command: a.Command + helpers.Option(noStdin)}
}

// SigProxy - Proxy all received signals to the process (default true)
func (a *Attach) SigProxy() *Attach {
	return &Attach{Command: a.Command + helpers.Option(sigProxy)}
}

func (a *Attach) ServiceName(serviceName ...string) *common.CommandExecutor {
	if len(serviceName) == 0 {
		return &common.CommandExecutor{Command: a.Command + helpers.ServiceName()}
	}
	return &common.CommandExecutor{Command: a.Command + helpers.ServiceName(serviceName[0])}
}
