package compose

import (
	"gocker/helpers"
)

type Build struct {
	Command string
}

// Options

// BuildArg - Set build-time variables for services
func (b *Build) BuildArg(arg, value string) *Build {
	b.Command += " --build-arg " + arg + "=" + value
	return b
}

// Builder - Set builder to use
func (b *Build) Builder(builderName string) *Build {
	b.Command += " --builder " + builderName
	return b
}

// DryRun - Execute command in dry run mode
func (b *Build) DryRun() *Build {
	b.Command += " --dry-run"
	return b
}

// Memory - Set memory limit for the build container. Not supported by BuildKit.
func (b *Build) Memory(bytes string) *Build {
	b.Command += " --memory " + bytes
	return b
}

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
