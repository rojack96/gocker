package common

import "github.com/rojack96/gocker/helpers"

const (
	format = "--format"
	index  = "--index"
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

func (ce *CommandExecutor) Exec() {
	helpers.GeneralExec(ce.Command, false)
}

func (ce *CommandExecutor) ExecWithPrivileges() {
	helpers.GeneralExec(ce.Command, true)
}

// Index - index of the container if service has multiple replicas
func Index(indexOfContainer int) string {
	return helpers.Int(index, indexOfContainer)
}

func Format(value string) string {
	return helpers.String(format, value)
}
