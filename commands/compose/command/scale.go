package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	scale = "--scale"
)

type Scale struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (s *Scale) DryRun() *Scale {
	return &Scale{Command: s.Command + option.DryRun()}
}

// NoDeps - Don't start linked services
func (s *Scale) NoDeps() *Scale {
	return &Scale{Command: s.Command + option.NoDeps()}
}

// ServiceNames - Specify services to scale (SERVICE=REPLICAS...)
func (s *Scale) ServiceNames(serviceReplicas ...helpers.KeyValueParameters) *common.CommandExecutor {
	var arguments []string

	if len(serviceReplicas) > 0 {
		for _, a := range serviceReplicas {
			arguments = append(arguments, a.Key+"="+a.Value)
		}
	}

	return &common.CommandExecutor{Command: s.Command + helpers.StringArray(scale, arguments...)}
}
