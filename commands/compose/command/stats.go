package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	noStream = "--no-stream"
)

type Stats struct {
	Command string
}

// All - Show all containers (including stopped ones)
func (s *Stats) All() *Stats {
	return &Stats{Command: s.Command + option.All()}
}

// DryRun - Execute command in dry run mode
func (s *Stats) DryRun() *Stats {
	return &Stats{Command: s.Command + option.DryRun()}
}

// Format - Format the output using a custom template
// 'table':            Print output in table format with column headers (default)
// 'table TEMPLATE':   Print output in table format using the given Go template
// 'json':             Print in JSON format
// 'TEMPLATE':         Print output using the given Go template.
func (s *Stats) Format(formatValue string) *Stats {
	return &Stats{Command: s.Command + common.Format(formatValue)}
}

// NoStream - Disable streaming stats and only pull the first result
func (s *Stats) NoStream() *Stats {
	return &Stats{Command: s.Command + helpers.Option(noStream)}
}

// NoTrunc - Do not truncate output
func (s *Stats) NoTrunc() *Stats {
	return &Stats{Command: s.Command + option.NoTrunc()}
}

func (s *Stats) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	if len(serviceNames) > 1 {
		return nil
	}
	return &common.CommandExecutor{Command: s.Command + helpers.ServiceName(serviceNames...)}
}
