package common

import (
	"github.com/rojack96/gocker/helpers"
	"strconv"
)

const (
	filter  = "--filter"
	format  = "--format"
	index   = "--index"
	pull    = "--pull"
	timeout = "--timeout"
	scale   = "--scale"
	env     = "--env"
	user    = "--user"
	workdir = "--workdir"
)

type Service interface {
	ServiceName(serviceName string) *CommandExecutor
}

type Services interface {
	ServiceNames(serviceNames ...string) *CommandExecutor
}

type CommandExecutor struct {
	command string
}

func SetCommand(cmd string) *CommandExecutor {
	return &CommandExecutor{command: cmd}
}

func (ce *CommandExecutor) GetCommand() string {
	return ce.command
}

/*
func (ce *CommandExecutor) GetCommandWithPrivileges() string {
	return ce.command
}
*/

func (ce *CommandExecutor) Exec(withPrivileges bool) {
	helpers.GeneralExec(ce.command, withPrivileges)
}

func Index(indexOfContainer int) string {
	return helpers.Int(index, indexOfContainer)
}

func Format(value string) string {
	return helpers.String(format, value)
}

func Pull(pullPolicy string) string {
	return helpers.String(pull, pullPolicy)
}

func Timeout(seconds int) string {
	return helpers.Int(timeout, seconds)
}

func Scale(service string, instances int) string {
	serviceScale := service + "=" + strconv.Itoa(instances)
	return helpers.String(scale, serviceScale)
}

func Filter(condition string) string {
	return helpers.String(filter, condition)
}

func Env(envs ...helpers.KeyValueParameters) string {
	var arguments []string

	if len(envs) > 0 {
		for _, a := range envs {
			arguments = append(arguments, a.Key+"="+a.Value)
		}
	}

	return helpers.StringArray(env, arguments...)
}

func User(usr string) string {
	return helpers.String(user, usr)
}

func Workdir(path string) string {
	return helpers.String(workdir, path)
}
