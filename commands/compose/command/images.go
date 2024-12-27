package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Images struct {
	command string
}

func NewImages(cmd string) *Images {
	return &Images{command: cmd}
}

// DryRun - Execute command in dry run mode
func (i *Images) DryRun() *Images {
	i.command += option.DryRun()
	return i
}

// Format - Format the output. Values: [table | json] (default "table")
func (i *Images) Format(value string) *Images {
	return &Images{command: i.command + common.Format(value)}
}

// Quiet - Only display IDs
func (i *Images) Quiet() *Images {
	return &Images{command: i.command + option.Quiet()}
}

func (i *Images) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(i.command + helpers.ServiceName(serviceNames...))
}
