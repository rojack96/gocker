package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
	"strings"
)

const (
	noTty      = "--no-TTY"
	privileged = "--privileged"
	user       = "--user"
	workdir    = "--workdir"
)

type Exec struct {
	Command string
}

// ToDo move to utils with build arg
type EnvArg struct {
	Key   string
	Value string
}

const (
	env = "--env"
)

/*
  -T, --no-TTY docker compose exec
      --privileged                   Give extended privileges to the process
  -u, --user string                  Run the command as this user
  -w, --workdir string               Path to workdir directory for this command
*/

// Detach - Detached mode: Run containers in the background
func (e *Exec) Detach() *Exec {
	return &Exec{Command: e.Command + option.Detach()}
}

// DryRun - Execute command in dry run mode
func (e *Exec) DryRun() *Exec {
	return &Exec{Command: e.Command + option.DryRun()}
}

// Env - Set environment variables
func (e *Exec) Env(envs ...EnvArg) *Exec {
	var arguments []string

	if len(envs) > 0 {
		for _, a := range envs {
			arguments = append(arguments, a.Key+"="+a.Value)
		}
	}

	return &Exec{Command: e.Command + helpers.StringArray(env, arguments...)}
}

// Index - Index of the container if service has multiple replicas
func (e *Exec) Index(indexOfContainer int) *Exec {
	return &Exec{Command: e.Command + common.Index(indexOfContainer)}
}

// NoTty - Disable pseudo-TTY allocation. By default docker compose exec allocates a TTY.
func (e *Exec) NoTty() *Exec {
	return &Exec{Command: e.Command + helpers.Option(noTty)}
}

// Privileged - Give extended privileges to the process
func (e *Exec) Privileged() *Exec {
	return &Exec{Command: e.Command + helpers.Option(privileged)}
}

// User - Run the command as this user
func (e *Exec) User(usr string) *Exec {
	return &Exec{Command: e.Command + helpers.String(user, usr)}
}

// Workdir - ath to workdir directory for this command
func (e *Exec) Workdir(path string) *Exec {
	return &Exec{Command: e.Command + helpers.String(workdir, path)}
}

func (e *Exec) ServiceCommandArgs(serviceName, command string, args ...string) *common.CommandExecutor {
	service := helpers.ServiceName(serviceName)
	cmd := service + " " + command

	if len(args) > 0 {
		cmd = cmd + " " + strings.Join(args, " ")
	}

	return &common.CommandExecutor{Command: e.Command + cmd}
}
