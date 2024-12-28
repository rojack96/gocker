package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

const (
	short = "--short"
)

type Version struct {
	command string
}

func NewVersion(cmd string) *Version {
	return &Version{command: cmd}
}

// DryRun - Execute command in dry run mode
func (v *Version) DryRun() *Version {
	return &Version{command: v.command + common.DryRun()}
}

// Format - Format the output. Values: [pretty | json]. (Default: pretty)
func (v *Version) Format(value string) *Version {
	return &Version{command: v.command + common.Format(value)}
}

// Short - Shows only Composes version number
func (v *Version) Short() *Version {
	return &Version{command: v.command + helpers.Option(short)}
}

func (v *Version) GetCommand() string {
	return v.command
}

func (v *Version) Exec(withPrivileges bool) {
	helpers.GeneralExec(v.command, withPrivileges)
}
