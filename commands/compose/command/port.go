package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
	"strconv"
)

const (
	protocol = "--protocol"
)

type Port struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (p *Port) DryRun() *Port {
	return &Port{Command: p.Command + option.DryRun()}
}

// Index - index of the container if service has multiple replicas
func (p *Port) Index(indexOfContainer int) *Port {
	return &Port{p.Command + common.Index(indexOfContainer)}
}

// Protocol - tcp or udp (default "tcp")
func (p *Port) Protocol(prtcl string) *Port {
	return &Port{p.Command + helpers.String(protocol, prtcl)}
}

func (e *Exec) ServiceAndPort(serviceName string, privatePort int) *common.CommandExecutor {
	service := helpers.ServiceName(serviceName)
	cmd := service + " " + strconv.Itoa(privatePort)

	return &common.CommandExecutor{Command: e.Command + cmd}
}
