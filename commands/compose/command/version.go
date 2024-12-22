package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	short = "--short"
)

type Version struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (v *Version) DryRun() *Version {
	return &Version{Command: v.Command + option.DryRun()}
}

// Format - Format the output. Values: [pretty | json]. (Default: pretty)
func (v *Version) Format(value string) *Version {
	return &Version{Command: v.Command + common.Format(value)}
}

// Short - Shows only Composes version number
func (v *Version) Short() *Version {
	return &Version{Command: v.Command + helpers.Option(short)}
}

func (v *Version) GetCommand() string {
	return v.Command
}

func (v *Version) Exec(withPrivileges bool) {
	helpers.GeneralExec(v.Command, withPrivileges)
}
