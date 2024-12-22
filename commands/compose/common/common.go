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

/*
func (ce *CommandExecutor) GetCommandWithPrivileges() string {
	return ce.Command
}
*/

func (ce *CommandExecutor) Exec(withPrivileges bool) {
	helpers.GeneralExec(ce.Command, withPrivileges)
}

// Index - index of the container if service has multiple replicas
func Index(indexOfContainer int) string {
	return helpers.Int(index, indexOfContainer)
}

func Format(value string) string {
	return helpers.String(format, value)
}
