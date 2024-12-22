package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

type Ls struct {
	Command string
}

// All - Show all stopped Compose projects
func (l *Ls) All() *Ls {
	return &Ls{Command: l.Command + option.All()}
}

// DryRun - Execute command in dry run mode
func (l *Ls) DryRun() *Ls {
	return &Ls{Command: l.Command + option.DryRun()}
}

// Filter - Filter output based on conditions provided
func (l *Ls) Filter(condition string) *Ls {
	return &Ls{Command: l.Command + common.Filter(condition)}
}

// Format - Format the output. Values: [table | json] (default "table")
func (l *Ls) Format(value string) *Ls {
	return &Ls{Command: l.Command + common.Format(value)}
}

// Quiet - Only display IDs
func (l *Ls) Quiet() *Ls {
	return &Ls{Command: l.Command + option.Quiet()}
}

func (l *Ls) GetCommand() string {
	return l.Command
}

func (l *Ls) Exec(withPrivileges bool) {
	helpers.GeneralExec(l.Command, withPrivileges)
}
