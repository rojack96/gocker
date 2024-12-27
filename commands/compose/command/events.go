package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	json = "--json"
)

type Events struct {
	command string
}

func NewEvents(cmd string) *Events {
	return &Events{command: cmd}
}

// DryRun - Execute command in dry run mode
func (e *Events) DryRun() *Events {
	return &Events{command: e.command + option.DryRun()}
}

// Json - Output events as a stream of json objects
func (e *Events) Json() *Events {
	return &Events{command: e.command + helpers.Option(json)}
}

func (e *Events) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: e.command + helpers.ServiceName(serviceNames...)}
}
