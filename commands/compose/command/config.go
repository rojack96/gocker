package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
)

const (
	hash                = "--hash"
	images              = "--images"
	noConsistency       = "--no-consistency"
	noInterpolate       = "--no-interpolate"
	noNormalize         = "--no-normalize"
	noPathResolution    = "--no-path-resolution"
	output              = "--output"
	profiles            = "--profiles"
	quiet               = "--quiet"
	resolveImageDigests = "--resolve-image-digests"
	variables           = "--variables"
)

type Config struct {
	command string
}

func NewConfig(cmd string) *Config {
	return &Config{command: cmd}
}

// DryRun - Execute command in dry run mode
func (c *Config) DryRun() *Config {
	return &Config{command: c.command + common.DryRun()}
}

// Format - Format the output. Values: [yaml | json] (default "yaml")
func (c *Config) Format(value string) *Config {
	return &Config{command: c.command + common.Format(value)}
}

// Hash - Print the service config hash, one per line
func (c *Config) Hash(value string) *Config {
	return &Config{command: c.command + helpers.String(hash, value)}
}

// Images - Print the image names, one per line
func (c *Config) Images() *Config {
	return &Config{command: c.command + helpers.Option(images)}
}

// NoConsistency - Don't check model consistency
func (c *Config) NoConsistency() *Config {
	return &Config{command: c.command + helpers.Option(noConsistency)}
}

// NoInterpolate - Don't interpolate environment variables
func (c *Config) NoInterpolate() *Config {
	return &Config{command: c.command + helpers.Option(noInterpolate)}
}

// NoNormalize - Don't normalize compose model
func (c *Config) NoNormalize() *Config {
	return &Config{command: c.command + helpers.Option(noNormalize)}
}

// NoPathResolution - Don't resolve file paths
func (c *Config) NoPathResolution() *Config {
	return &Config{command: c.command + helpers.Option(noPathResolution)}
}

// Output - Save to file (default to stdout)
func (c *Config) Output(value string) *Config {
	return &Config{command: c.command + helpers.String(output, value)}
}

// Profiles - Print the profile names, one per line
func (c *Config) Profiles() *Config {
	return &Config{command: c.command + helpers.Option(profiles)}
}

// Quiet - Only validate the configuration, don't print anything
func (c *Config) Quiet() *Config {
	return &Config{command: c.command + helpers.Option(quiet)}
}

// ResolveImageDigests - Pin image tags to digests
func (c *Config) ResolveImageDigests() *Config {
	return &Config{command: c.command + helpers.Option(resolveImageDigests)}
}

// Services - Print the service names, one per line
func (c *Config) Services() *Config {
	return &Config{command: c.command + common.ServicesOption()}
}

// Variables - Print model variables and default values
func (c *Config) Variables() *Config {
	return &Config{command: c.command + helpers.Option(variables)}
}

// Volumes - Print the volume names, one per line
func (c *Config) Volumes() *Config {
	return &Config{command: c.command + common.Volumes()}
}

func (c *Config) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(c.command + helpers.ServiceName(serviceNames...))
}
