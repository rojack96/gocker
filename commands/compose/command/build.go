package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	buildArg         = "--build-arg"
	builder          = "--builder"
	memory           = "--memory"
	noCache          = "--no-cache"
	pull             = "--pull"
	push             = "--push"
	ssh              = "--ssh"
	withDependencies = "--with-dependencies"
)

type Build struct {
	command string
}

func NewBuild(cmd string) *Build {
	return &Build{command: cmd}
}

type UnitByte string

const (
	Kilobytes UnitByte = "k"
	Megabytes UnitByte = "m"
	Gigabytes UnitByte = "g"
)

// Options

// BuildArg - Set command-time variables for services
func (b *Build) BuildArg(args ...helpers.KeyValueParameters) *Build {
	var arguments []string

	if len(args) > 0 {
		for _, a := range args {
			arguments = append(arguments, a.Key+"="+a.Value)
		}
	}

	return &Build{command: b.command + helpers.StringArray(buildArg, arguments...)}
}

// Builder - Set builder to use
func (b *Build) Builder(builderName string) *Build {
	return &Build{command: b.command + helpers.String(builder, builderName)}
}

// DryRun - Execute command in dry run mode
func (b *Build) DryRun() *Build {
	return &Build{command: b.command + option.DryRun()}
}

// Memory - Set memory limit for the command container. Not supported by BuildKit.
func (b *Build) Memory(bytes string, unitByte UnitByte) *Build {
	bytes += string(unitByte)
	return &Build{command: b.command + helpers.String(memory, bytes)}
}

// NoCache - Do not use cache when building the image
func (b *Build) NoCache() *Build {
	return &Build{command: b.command + helpers.Option(noCache)}
}

// Pull - Always attempt to pull a newer version of the image
func (b *Build) Pull() *Build {
	return &Build{command: b.command + helpers.Option(pull)}
}

// Push - Push service images
func (b *Build) Push() *Build {
	return &Build{command: b.command + helpers.Option(push)}
}

// Quiet - Don't print anything to STDOUT
func (b *Build) Quiet() *Build {
	return &Build{command: b.command + option.Quiet()}
}

// Ssh - Set SSH authentications used when building service images.
// (use 'default' for using your default SSH Agent)
func (b *Build) Ssh(agent string) *Build {
	return &Build{command: b.command + helpers.String(ssh, agent)}
}

// WithDependencies - Also command dependencies (transitively)
func (b *Build) WithDependencies() *Build {
	return &Build{command: b.command + helpers.Option(withDependencies)}
}

func (b *Build) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: b.command + helpers.ServiceName(serviceNames...)}
}
