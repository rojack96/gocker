package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	rim = "--rmi"
)

type Down struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (d *Down) DryRun() *Down {
	return &Down{Command: d.Command + option.DryRun()}
}

// RemoveOrphans - Remove containers for services not defined in the Compose file
func (d *Down) RemoveOrphans() *Down {
	return &Down{Command: d.Command + option.RemoveOrphans()}
}

// Rmi - Remove images used by services. "local" remove only images that don't have a custom tag ("local"|"all")
func (d *Down) Rmi(value string) *Down {
	return &Down{Command: d.Command + helpers.String(rim, value)}
}

func (d *Down) Timeout(seconds int) *Down {
	return &Down{Command: d.Command + common.Timeout(seconds)}
}

// Volumes - Remove named volumes declared in the "volumes" section of the Compose file and anonymous volumes attached to containers
func (d *Down) Volumes() *Down {
	return &Down{Command: d.Command + option.Volumes()}
}
