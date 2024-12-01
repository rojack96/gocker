package compose

import (
	"gocker/helpers"
)

type Build struct {
	Command string
}

func (b *Build) NoCache() {
	b.Command += " --no-cache"
}

func (b *Build) ServiceName(serviceName ...string) *Build {
	if len(serviceName) != 0 {
		for _, service := range serviceName {
			b.Command += helpers.AppendString(service)
		}
	}

	return b
}

func (b *Build) GetCommand() string {
	return b.Command
}

func (b *Build) Exec() {
	helpers.GeneralExec(b.Command)
}
