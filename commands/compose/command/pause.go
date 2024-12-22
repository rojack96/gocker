package command

import "github.com/rojack96/gocker/commands/compose/option"

type Pause struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (p *Pause) DryRun() *Pause {
	return &Pause{Command: p.Command + option.DryRun()}
}
