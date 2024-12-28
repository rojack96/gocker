package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
	"github.com/rojack96/gocker/options"
	"strings"
)

const (
	capAdd       = "--cap-add"
	capDrop      = "--cap-drop"
	entrypoint   = "--entrypoint"
	interactive  = "--interactive"
	label        = "--label"
	name         = "--name"
	publish      = "--publish"
	rm           = "--rm"
	servicePorts = "--service-ports"
	useAliases   = "--use-aliases"
	volume       = "--volume"
)

type Run struct {
	command string
}

func NewRun(cmd string) *Run {
	return &Run{command: cmd}
}

// Build - Build image before starting container
func (r *Run) Build() *Run {
	return &Run{command: r.command + options.Build()}
}

// CapAdd - Add Linux capabilities
func (r *Run) CapAdd(capabilities ...string) *Run {
	return &Run{command: r.command + helpers.StringArray(capAdd, capabilities...)}
}

// CapDrop - Drop Linux capabilities
func (r *Run) CapDrop(capabilities ...string) *Run {
	return &Run{command: r.command + helpers.StringArray(capDrop, capabilities...)}
}

// Detach - Run container in background and print container ID
func (r *Run) Detach() *Run {
	return &Run{command: r.command + options.Detach()}
}

// DryRun - Execute command in dry run mode
func (r *Run) DryRun() *Run {
	return &Run{command: r.command + options.DryRun()}
}

// Entrypoint - Override the entrypoint of the image
func (r *Run) Entrypoint(value string) *Run {
	return &Run{command: r.command + helpers.String(entrypoint, value)}
}

// Env - Set environment variables
func (r *Run) Env(envs ...helpers.KeyValueParameters) *Run {
	return &Run{command: r.command + common.Env(envs...)}
}

// Interactive - Keep STDIN open even if not attached (default true)
func (r *Run) Interactive() *Run {
	return &Run{command: r.command + helpers.Option(interactive)}
}

// Label - Add or override a label
func (r *Run) Label(labels ...string) *Run {
	return &Run{command: r.command + helpers.StringArray(label, labels...)}
}

// Name - Assign a name to the container
func (r *Run) Name(value string) *Run {
	return &Run{command: r.command + helpers.String(name, value)}
}

// NoTTY - Disable pseudo-TTY allocation (default: auto-detected)
func (r *Run) NoTTY() *Run {
	return &Run{command: r.command + options.NoTty()}
}

// NoDeps - Don't start linked services
func (r *Run) NoDeps() *Run {
	return &Run{command: r.command + options.NoDeps()}
}

// Publish - Publish a container's port(s) to the host
func (r *Run) Publish(ports ...string) *Run {
	return &Run{command: r.command + helpers.StringArray(publish, ports...)}
}

// QuietPull - Pull without printing progress information
func (r *Run) QuietPull() *Run {
	return &Run{command: r.command + options.QuietPull()}
}

// RemoveOrphans - Remove containers for services not defined in the Compose file
func (r *Run) RemoveOrphans() *Run {
	return &Run{command: r.command + options.RemoveOrphans()}
}

// RM - Automatically remove the container when it exits
func (r *Run) RM() *Run {
	return &Run{command: r.command + helpers.Option(rm)}
}

// ServicePorts - Run command with all service's ports enabled and mapped to the host
func (r *Run) ServicePorts() *Run {
	return &Run{command: r.command + helpers.Option(servicePorts)}
}

// UseAliases - Use the service's network useAliases in the network(s) the container connects to
func (r *Run) UseAliases() *Run {
	return &Run{command: r.command + helpers.Option(useAliases)}
}

// User - Run as specified username or uid
func (r *Run) User(value string) *Run {
	return &Run{command: r.command + common.User(value)}
}

// Volume - Bind mount a volume
func (r *Run) Volume(volumes ...string) *Run {
	return &Run{command: r.command + helpers.StringArray(volume, volumes...)}
}

// Workdir - Working directory inside the container
func (r *Run) Workdir(value string) *Run {
	return &Run{command: r.command + common.Workdir(value)}
}

// ServiceCommandArgs - Specify services to up
func (r *Run) ServiceCommandArgs(serviceNames string, command *string, args ...string) *common.CommandExecutor {
	service := helpers.ServiceName(serviceNames)
	cmd := service

	if command != nil {
		cmd = cmd + " " + *command
		if len(args) > 0 {
			cmd = cmd + " " + strings.Join(args, " ")
		}
	}

	return common.SetCommand(r.command + helpers.ServiceName(serviceNames))
}
