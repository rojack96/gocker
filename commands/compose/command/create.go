package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Create struct {
	command string
}

func NewCreate(cmd string) *Create {
	return &Create{command: cmd}
}

// Build - Build images before starting containers
func (c *Create) Build() *Create {
	return &Create{command: c.command + option.Build()}
}

// DryRun - Execute command in dry run mode
func (c *Create) DryRun() *Create {
	return &Create{command: c.command + option.DryRun()}
}

// ForceRecreate - Recreate containers even if their configuration and image haven't changed
func (c *Create) ForceRecreate() *Create {
	return &Create{command: c.command + option.ForceRecreate()}
}

// NoBuild - Don't build an image, even if it's policy
func (c *Create) NoBuild() *Create {
	return &Create{command: c.command + option.NoBuild()}
}

// NoRecreate - If containers already exist, don't recreate them. Incompatible with --force-recreate.
func (c *Create) NoRecreate() *Create {
	return &Create{command: c.command + option.NoRecreate()}
}

// Pull - Pull image before running ("always"|"missing"|"never"|"build") (default "policy")
func (c *Create) Pull(pullPolicy string) *Create {
	return &Create{command: c.command + common.Pull(pullPolicy)}
}

// QuietPull - Pull without printing progress information
func (c *Create) QuietPull() *Create {
	return &Create{command: c.command + option.QuietPull()}
}

// RemoveOrphans - Remove containers for services not defined in the Compose file
func (c *Create) RemoveOrphans() *Create {
	return &Create{command: c.command + option.RemoveOrphans()}
}

// Scale - Scale SERVICE to NUM instances. Overrides the scale setting in the Compose file if present.
func (c *Create) Scale(service string, instances int) *Create {
	return &Create{command: c.command + common.Scale(service, instances)}
}

func (c *Create) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(c.command + helpers.ServiceName(serviceNames...))
}
