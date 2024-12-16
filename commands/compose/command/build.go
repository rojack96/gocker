package command

import (
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
	quiet            = "--quiet"
	ssh              = "--ssh"
	withDependencies = "--with-dependencies"
)

type Build struct {
	Command string
}

type BuildArg struct {
	Key   string
	Value string
}

type UnitByte string

const (
	Kilobytes UnitByte = "k"
	Megabytes UnitByte = "m"
	Gigabytes UnitByte = "g"
)

// Options

// BuildArg - Set command-time variables for services
func (b *Build) BuildArg(args ...BuildArg) *Build {
	var arguments []string

	if len(args) > 0 {
		for _, a := range args {
			arguments = append(arguments, a.Key+"="+a.Value)
		}
	}

	b.Command += helpers.StringArray(buildArg, arguments...)
	return b
}

// Builder - Set builder to use
func (b *Build) Builder(builderName string) *Build {
	b.Command += helpers.String(builder, builderName)
	return b
}

// DryRun - Execute command in dry run mode
func (b *Build) DryRun() *Build {
	b.Command += option.DryRun()
	return b
}

// Memory - Set memory limit for the command container. Not supported by BuildKit.
func (b *Build) Memory(bytes string, unitByte UnitByte) *Build {
	bytes += string(unitByte)
	b.Command += helpers.String(memory, bytes)
	return b
}

// TODO continue from here

// NoCache - Do not use cache when building the image
func (b *Build) NoCache() *Build {
	b.Command += helpers.Option(noCache)
	return b
}

// Pull - Always attempt to pull a newer version of the image
func (b *Build) Pull() *Build {
	b.Command += helpers.Option(pull)
	return b
}

// Push - Push service images
func (b *Build) Push() *Build {
	b.Command += helpers.Option(push)
	return b
}

// Quiet - Don't print anything to STDOUT
func (b *Build) Quiet() *Build {
	b.Command += helpers.Option(quiet)
	return b
}

// Ssh - Set SSH authentications used when building service images.
// (use 'default' for using your default SSH Agent)
func (b *Build) Ssh(agent string) *Build {
	b.Command += helpers.String(ssh, agent)
	return b
}

// WithDependencies - Also command dependencies (transitively)
func (b *Build) WithDependencies() *Build {
	b.Command += helpers.Option(withDependencies)
	return b
}

func (b *Build) ServiceName(serviceName ...string) *Build {
	b.Command = helpers.ServiceName(serviceName...)
	return b
}

func (b *Build) GetCommand() string {
	return b.Command
}

func (b *Build) Exec() {
	helpers.GeneralExec(b.Command)
}
