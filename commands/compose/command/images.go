package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	format = "--format"
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
	i.Command += common.Format(value)
	return i
}

// Quiet - Only display IDs
func (i *Images) Quiet() *Images {
	i.Command += option.Quiet()
	return i
}

func (i *Images) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	i.Command += helpers.ServiceName(serviceNames...)
	return &common.CommandExecutor{Command: i.Command}
}
