package common

import (
	"github.com/rojack96/gocker/helpers"
	"strconv"
)

const (
	format  = "--format"
	index   = "--index"
	pull    = "--pull"
	timeout = "--timeout"
	scale   = "--scale"
)

type Service interface {
	ServiceName(serviceName string) *CommandExecutor
}

type Services interface {
	ServiceNames(serviceNames ...string) *CommandExecutor
}

type CommandExecutor struct {
	Command string
}

func (ce *CommandExecutor) GetCommand() string {
	return ce.Command
}

/*
func (ce *CommandExecutor) GetCommandWithPrivileges() string {
	return ce.Command
}
*/

func (ce *CommandExecutor) Exec(withPrivileges bool) {
	helpers.GeneralExec(ce.Command, withPrivileges)
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
