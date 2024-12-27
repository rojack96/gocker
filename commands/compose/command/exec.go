package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
	"strings"
)

const (
	privileged = "--privileged"
)

type Exec struct {
	command string
}

func NewExec(cmd string) *Exec {
	return &Exec{command: cmd}
}

// Detach - Detached mode: Run containers in the background
func (e *Exec) Detach() *Exec {
	return &Exec{command: e.command + option.Detach()}
}

// DryRun - Execute command in dry run mode
func (e *Exec) DryRun() *Exec {
	return &Exec{command: e.command + option.DryRun()}
}

// Env - Set environment variables
func (e *Exec) Env(envs ...helpers.KeyValueParameters) *Exec {
	return &Exec{command: e.command + common.Env(envs...)}
}

// Index - Index of the container if service has multiple replicas
func (e *Exec) Index(indexOfContainer int) *Exec {
	return &Exec{command: e.command + common.Index(indexOfContainer)}
}

// NoTty - Disable pseudo-TTY allocation. By default, docker compose exec allocates a TTY.
func (e *Exec) NoTty() *Exec {
	return &Exec{command: e.command + option.NoTty()}
}

// Privileged - Give extended privileges to the process
func (e *Exec) Privileged() *Exec {
	return &Exec{command: e.command + helpers.Option(privileged)}
}

// User - Run the command as this user
func (e *Exec) User(usr string) *Exec {
	return &Exec{command: e.command + common.User(usr)}
}

// Workdir - ath to workdir directory for this command
func (e *Exec) Workdir(path string) *Exec {
	return &Exec{command: e.command + common.Workdir(path)}
}

func (e *Exec) ServiceCommandArgs(serviceName, command string, args ...string) *common.CommandExecutor {
	service := helpers.ServiceName(serviceName)
	cmd := service + " " + command

	if len(args) > 0 {
		cmd = cmd + " " + strings.Join(args, " ")
	}

	return common.SetCommand(e.command + cmd)
}
