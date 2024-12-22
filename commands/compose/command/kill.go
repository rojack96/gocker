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
	Command string
}

// DryRun - Execute command in dry run mode
func (k *Kill) DryRun() *Kill {
	return &Kill{Command: k.Command + option.DryRun()}
}

// RemoveOrphans - Remove containers for services not defined in the Compose file
func (k *Kill) RemoveOrphans() *Kill {
	return &Kill{Command: k.Command + option.RemoveOrphans()}
}

func (k *Kill) Signal(value string) *Kill {
	return &Kill{Command: k.Command + helpers.String(signal, value)}
}

func (k *Kill) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: k.Command + helpers.ServiceName(serviceNames...)}
}
