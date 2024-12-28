package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

const (
	noStream = "--no-stream"
)

type Stats struct {
	command string
}

func NewStats(cmd string) *Stats {
	return &Stats{command: cmd}
}

// All - Show all containers (including stopped ones)
func (s *Stats) All() *Stats {
	return &Stats{command: s.command + common.All()}
}

// DryRun - Execute command in dry run mode
func (s *Stats) DryRun() *Stats {
	return &Stats{command: s.command + common.DryRun()}
}

// Format - Format the output using a custom template
// 'table':            Print output in table format with column headers (default)
// 'table TEMPLATE':   Print output in table format using the given Go template
// 'json':             Print in JSON format
// 'TEMPLATE':         Print output using the given Go template.
func (s *Stats) Format(formatValue string) *Stats {
	return &Stats{command: s.command + common.Format(formatValue)}
}

// NoStream - Disable streaming stats and only pull the first result
func (s *Stats) NoStream() *Stats {
	return &Stats{command: s.command + helpers.Option(noStream)}
}

// NoTrunc - Do not truncate output
func (s *Stats) NoTrunc() *Stats {
	return &Stats{command: s.command + common.NoTrunc()}
}

func (s *Stats) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	if len(serviceNames) > 1 {
		return nil
	}
	return common.SetCommand(s.command + helpers.ServiceName(serviceNames...))
}
