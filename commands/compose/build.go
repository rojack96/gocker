package compose

import (
	"gocker/helpers"
)

const (
	buildArg = "--build-arg"
	builder  = "--builder"
	memory   = "--memory"
	/*
	 Options:
	       --build-arg stringArray   Set build-time variables for services
	       --builder string          Set builder to use
	       --dry-run                 Execute command in dry run mode
	   -m, --memory bytes            Set memory limit for the build container. Not supported by BuildKit.
	       --no-cache                Do not use cache when building the image
	       --pull                    Always attempt to pull a newer version of the image
	       --push                    Push service images
	   -q, --quiet                   Don't print anything to STDOUT
	       --ssh string              Set SSH authentications used when building service images. (use 'default' for using your default SSH Agent)
	       --with-dependencies       Also build dependencies (transitively)
	*/
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

// BuildArg - Set build-time variables for services
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
	b.Command += helpers.Option(dryRun)
	return b
}

// Memory - Set memory limit for the build container. Not supported by BuildKit.
func (b *Build) Memory(bytes string, unitByte UnitByte) *Build {
	bytes += string(unitByte)
	b.Command += helpers.String(memory, bytes)
	return b
}

// TODO continue from here

// NoCache - Do not use cache when building the image
func (b *Build) NoCache() *Build {
	b.Command += " --no-cache"
	return b
}

// Pull - Always attempt to pull a newer version of the image
func (b *Build) Pull() *Build {
	b.Command += " --pull"
	return b
}

// Push - Push service images
func (b *Build) Push() *Build {
	b.Command += " --push"
	return b
}

// Quiet - Don't print anything to STDOUT
func (b *Build) Quiet() *Build {
	b.Command += " --quiet"
	return b
}

// Ssh - Set SSH authentications used when building service images.
// (use 'default' for using your default SSH Agent)
func (b *Build) Ssh(agent string) *Build {
	b.Command += " --ssh " + agent
	return b
}

// WithDependencies - Also build dependencies (transitively)
func (b *Build) WithDependencies() *Build {
	b.Command += " --with-dependencies"
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
