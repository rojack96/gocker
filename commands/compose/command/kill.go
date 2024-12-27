package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	signal = "--signal"
)

type Kill struct {
	command string
}

func NewKill(cmd string) *Kill {
	return &Kill{command: cmd}
}

// DryRun - Execute command in dry run mode
func (k *Kill) DryRun() *Kill {
	return &Kill{command: k.command + option.DryRun()}
}

// RemoveOrphans - Remove containers for services not defined in the Compose file
func (k *Kill) RemoveOrphans() *Kill {
	return &Kill{command: k.command + option.RemoveOrphans()}
}

func (k *Kill) Signal(value string) *Kill {
	return &Kill{command: k.command + helpers.String(signal, value)}
}

func (k *Kill) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(k.command + helpers.ServiceName(serviceNames...))
}
