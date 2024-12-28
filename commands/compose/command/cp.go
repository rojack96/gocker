package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

const (
	archive    = "--archive"
	followLink = "--follow-link"
)

type Cp struct {
	command string
}

func NewCp(cmd string) *Cp {
	return &Cp{command: cmd}
}

type Copy struct {
	Reverse     bool
	Service     string
	LocalFile   string
	ServiceFile string
}

// Archive - Archive mode (copy all uid/gid information)
func (cp *Cp) Archive() *Cp {
	return &Cp{command: cp.command + helpers.Option(archive)}
}

// DryRun - Execute command in dry run mode
func (cp *Cp) DryRun() *Cp {
	return &Cp{command: cp.command + common.DryRun()}
}

// FollowLink - Always follow symbol link in SRC_PATH
func (cp *Cp) FollowLink() *Cp {
	return &Cp{command: cp.command + helpers.Option(followLink)}
}

// Index - Index of the container if service has multiple replicas
func (cp *Cp) Index(idx int) *Cp {
	return &Cp{command: cp.command + common.Index(idx)}
}

func (cp *Cp) Copy(copy Copy) *common.CommandExecutor {
	localToService := copy.LocalFile + " " + copy.Service + ":" + copy.ServiceFile
	serviceToLocal := copy.Service + ":" + copy.ServiceFile + " " + copy.LocalFile

	if copy.Reverse {
		return common.SetCommand(cp.command + " " + serviceToLocal)
	}
	return common.SetCommand(cp.command + " " + localToService)
}
