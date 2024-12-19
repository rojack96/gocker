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
	Command string
}

// DryRun - Execute command in dry run mode
func (l *Logs) DryRun() *Logs {
	l.Command += option.DryRun()
	return l
}

// Follow - Follow log output
func (l *Logs) Follow() *Logs {
	l.Command += helpers.Option(follow)
	return l
}

// Index - index of the container if service has multiple replicas
func (l *Logs) Index(indexOfContainer int) *Logs {
	l.Command += common.Index(indexOfContainer)
	return l
}

// NoColor - Produce monochrome output
func (l *Logs) NoColor() *Logs {
	l.Command += option.NoColor()
	return l
}

// NoLogPrefix - Don't print prefix in logs
func (l *Logs) NoLogPrefix() *Logs {
	l.Command += option.NoLogPrefix()
	return l
}

// Timestamps - Show timestamps
func (l *Logs) Timestamps() *Logs {
	l.Command += option.Timestamps()
	return l
}

// Tail - Number of lines to show from the end of the logs for each container (default "all")
func (l *Logs) Tail(value string) *Logs {
	l.Command += helpers.String(tail, value)
	return l
}

// Since - Show logs since timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)
func (l *Logs) Since(date string) *Logs {
	l.Command += helpers.String(since, date)
	return l
}

// Until - Show logs before a timestamp (e.g. 2013-01-02T13:23:37Z) or relative (e.g. 42m for 42 minutes)
func (l *Logs) Until(date string) *Logs {
	l.Command += helpers.String(until, date)
	return l
}

func (l *Logs) GetCommand() string {
	return l.Command
}

func (l *Logs) Exec() {
	helpers.GeneralExec(l.Command, false)
}

func (l *Logs) ExecWithPrivileges() {
	helpers.GeneralExec(l.Command, true)
}
