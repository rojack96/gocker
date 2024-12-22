package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Images struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (i *Images) DryRun() *Images {
	i.Command += option.DryRun()
	return i
}

// Format - Format the output. Values: [table | json] (default "table")
func (i *Images) Format(value string) *Images {
	return &Images{Command: i.Command + common.Format(value)}
}

// Quiet - Only display IDs
func (i *Images) Quiet() *Images {
	return &Images{Command: i.Command + option.Quiet()}
}

func (i *Images) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: i.Command + helpers.ServiceName(serviceNames...)}
}
