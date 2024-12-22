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
	Command string
}

// DryRun - Execute command in dry run mode
func (e *Events) DryRun() *Events {
	return &Events{Command: e.Command + option.DryRun()}
}

// Json - Output events as a stream of json objects
func (e *Events) Json() *Events {
	return &Events{Command: e.Command + helpers.Option(json)}
}

func (e *Events) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: e.Command + helpers.ServiceName(serviceNames...)}
}
