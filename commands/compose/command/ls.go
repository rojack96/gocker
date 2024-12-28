package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

type Ls struct {
	command string
}

func NewLs(cmd string) *Ls {
	return &Ls{command: cmd}
}

// All - Show all stopped Compose projects
func (l *Ls) All() *Ls {
	return &Ls{command: l.command + common.All()}
}

// DryRun - Execute command in dry run mode
func (l *Ls) DryRun() *Ls {
	return &Ls{command: l.command + common.DryRun()}
}

// Filter - Filter output based on conditions provided
func (l *Ls) Filter(condition string) *Ls {
	return &Ls{command: l.command + common.Filter(condition)}
}

// Format - Format the output. Values: [table | json] (default "table")
func (l *Ls) Format(value string) *Ls {
	return &Ls{command: l.command + common.Format(value)}
}

// Quiet - Only display IDs
func (l *Ls) Quiet() *Ls {
	return &Ls{command: l.command + common.Quiet()}
}

func (l *Ls) GetCommand() string {
	return l.command
}

func (l *Ls) Exec(withPrivileges bool) {
	helpers.GeneralExec(l.command, withPrivileges)
}
