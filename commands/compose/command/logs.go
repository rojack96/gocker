package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	follow = "--follow"
	since  = "--since"
	tail   = "--tail"
	until  = "--until"
)

type Logs struct {
	command string
}

func NewLogs(cmd string) *Logs {
	return &Logs{command: cmd}
}

// DryRun - Execute command in dry run mode
func (l *Logs) DryRun() *Logs {
	return &Logs{command: l.command + option.DryRun()}
}

// Follow - Follow log output
func (l *Logs) Follow() *Logs {
	return &Logs{command: l.command + helpers.Option(follow)}
}

// Index - index of the container if service has multiple replicas
func (l *Logs) Index(indexOfContainer int) *Logs {
	return &Logs{command: l.command + common.Index(indexOfContainer)}
}

// NoColor - Produce monochrome output
func (l *Logs) NoColor() *Logs {
	return &Logs{command: l.command + option.NoColor()}
}

// NoLogPrefix - Don't print prefix in logs
func (l *Logs) NoLogPrefix() *Logs {
	return &Logs{command: l.command + option.NoLogPrefix()}
}

// Timestamps - Show timestamps
func (l *Logs) Timestamps() *Logs {
	return &Logs{command: l.command + option.Timestamps()}
}

// Tail - Number of lines to show from the end of the logs for each container (default "all")
func (l *Logs) Tail(value string) *Logs {
	return &Logs{command: l.command + helpers.String(tail, value)}
}

// Since - Show logs since timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)
func (l *Logs) Since(date string) *Logs {
	return &Logs{command: l.command + helpers.String(since, date)}
}

// Until - Show logs before a timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)
func (l *Logs) Until(date string) *Logs {
	return &Logs{command: l.command + helpers.String(until, date)}
}

func (l *Logs) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: l.command + helpers.ServiceName(serviceNames...)}
}
