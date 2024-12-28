package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

const (
	rim = "--rmi"
)

type Down struct {
	command string
}

func NewDown(cmd string) *Down {
	return &Down{command: cmd}
}

// DryRun - Execute command in dry run mode
func (d *Down) DryRun() *Down {
	return &Down{command: d.command + common.DryRun()}
}

// RemoveOrphans - Remove containers for services not defined in the Compose file
func (d *Down) RemoveOrphans() *Down {
	return &Down{command: d.command + common.RemoveOrphans()}
}

// Rmi - Remove images used by services. "local" remove only images that don't have a custom tag ("local"|"all")
func (d *Down) Rmi(value string) *Down {
	return &Down{command: d.command + helpers.String(rim, value)}
}

func (d *Down) Timeout(seconds int) *Down {
	return &Down{command: d.command + common.Timeout(seconds)}
}

// Volumes - Remove named volumes declared in the "volumes" section of the Compose file and anonymous volumes attached to containers
func (d *Down) Volumes() *Down {
	return &Down{command: d.command + common.Volumes()}
}

func (d *Down) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	if len(serviceNames) > 1 {
		return nil
	}
	return common.SetCommand(d.command + helpers.ServiceName(serviceNames...))
}
