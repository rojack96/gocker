package compose

import (
	"gocker/helpers"
)

type Up struct {
	Command string
}

func (up *Up) Build() *Up {
	up.Command += " --build"
	return up
}

func (up *Up) Detach() *Up {
	up.Command += " -d"
	return up
}

func (up *Up) ServiceName(serviceName ...string) *Up {
	if len(serviceName) != 0 {
		for _, service := range serviceName {
			up.Command += helpers.AppendString(service)
		}
	}

	return up
}

func (up *Up) GetCommand() string {
	return up.Command
}

func (up *Up) Exec() {
	helpers.GeneralExec(up.Command)
}
