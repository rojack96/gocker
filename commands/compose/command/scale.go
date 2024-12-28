package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

const (
	scale = "--scale"
)

type Scale struct {
	command string
}

func NewScale(cmd string) *Scale {
	return &Scale{command: cmd}
}

// DryRun - Execute command in dry run mode
func (s *Scale) DryRun() *Scale {
	return &Scale{command: s.command + common.DryRun()}
}

// NoDeps - Don't start linked services
func (s *Scale) NoDeps() *Scale {
	return &Scale{command: s.command + common.NoDeps()}
}

// ServiceNames - Specify services to scale (SERVICE=REPLICAS...)
func (s *Scale) ServiceNames(serviceReplicas ...helpers.KeyValueParameters) *common.CommandExecutor {
	var arguments []string

	if len(serviceReplicas) > 0 {
		for _, a := range serviceReplicas {
			arguments = append(arguments, a.Key+"="+a.Value)
		}
	}

	return common.SetCommand(s.command + helpers.StringArray(scale, arguments...))
}
