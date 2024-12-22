package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	archive    = "--archive"
	followLink = "--follow-link"
)

type Cp struct {
	Command string
}

type Copy struct {
	Reverse     bool
	Service     string
	LocalFile   string
	ServiceFile string
}

// Archive - Archive mode (copy all uid/gid information)
func (cp *Cp) Archive() *Cp {
	return &Cp{Command: cp.Command + helpers.Option(archive)}
}

// DryRun - Execute command in dry run mode
func (cp *Cp) DryRun() *Cp {
	return &Cp{Command: cp.Command + option.DryRun()}
}

// FollowLink - Always follow symbol link in SRC_PATH
func (cp *Cp) FollowLink() *Cp {
	return &Cp{Command: cp.Command + helpers.Option(followLink)}
}

// Index - Index of the container if service has multiple replicas
func (cp *Cp) Index(idx int) *Cp {
	return &Cp{Command: cp.Command + common.Index(idx)}
}

func (cp *Cp) Copy(copy Copy) *common.CommandExecutor {
	localToService := copy.LocalFile + " " + copy.Service + ":" + copy.ServiceFile
	serviceToLocal := copy.Service + ":" + copy.ServiceFile + " " + copy.LocalFile

	if copy.Reverse {
		return &common.CommandExecutor{Command: cp.Command + " " + serviceToLocal}
	}

	return &common.CommandExecutor{Command: cp.Command + " " + localToService}
}
