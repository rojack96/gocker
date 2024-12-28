package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
	"strconv"
)

const (
	protocol = "--protocol"
)

type Port struct {
	command string
}

func NewPort(cmd string) *Port {
	return &Port{command: cmd}
}

// DryRun - Execute command in dry run mode
func (p *Port) DryRun() *Port {
	return &Port{command: p.command + common.DryRun()}
}

// Index - index of the container if service has multiple replicas
func (p *Port) Index(indexOfContainer int) *Port {
	return &Port{p.command + common.Index(indexOfContainer)}
}

// Protocol - tcp or udp (default "tcp")
func (p *Port) Protocol(value string) *Port {
	return &Port{p.command + helpers.String(protocol, value)}
}

func (p *Port) ServiceAndPort(serviceName string, privatePort int) *common.CommandExecutor {
	service := helpers.ServiceName(serviceName)
	cmd := service + " " + strconv.Itoa(privatePort)

	return common.SetCommand(p.command + cmd)
}
