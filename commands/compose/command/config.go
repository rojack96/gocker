package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
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
	services            = "--services"
	variables           = "--variables"
	volumes             = "--volumes"
)

type Config struct {
	Command string
}

// DryRun - Execute command in dry run mode
func (c *Config) DryRun() *Config {
	c.Command += option.DryRun()
	return c
}

// Format - Format the output. Values: [yaml | json] (default "yaml")
func (c *Config) Format(value string) *Config {
	c.Command += common.Format(value)
	return c
}

// Hash - Print the service config hash, one per line
func (c *Config) Hash(value string) *Config {
	c.Command += helpers.String(hash, value)
	return c
}

// Images - Print the image names, one per line
func (c *Config) Images() *Config {
	c.Command += helpers.Option(images)
	return c
}

// NoConsistency - Don't check model consistency
func (c *Config) NoConsistency() *Config {
	c.Command += helpers.Option(noConsistency)
	return c
}

// NoInterpolate - Don't interpolate environment variables
func (c *Config) NoInterpolate() *Config {
	c.Command += helpers.Option(noInterpolate)
	return c
}

// NoNormalize - Don't normalize compose model
func (c *Config) NoNormalize() *Config {
	c.Command += helpers.Option(noNormalize)
	return c
}

// NoPathResolution - Don't resolve file paths
func (c *Config) NoPathResolution() *Config {
	c.Command += helpers.Option(noPathResolution)
	return c
}

// Output - Save to file (default to stdout)
func (c *Config) Output(value string) *Config {
	c.Command += helpers.String(output, value)
	return c
}

// Profiles - Print the profile names, one per line
func (c *Config) Profiles() *Config {
	c.Command += helpers.Option(profiles)
	return c
}

// Quiet - Only validate the configuration, don't print anything
func (c *Config) Quiet() *Config {
	c.Command += helpers.Option(quiet)
	return c
}

// ResolveImageDigests - Pin image tags to digests
func (c *Config) ResolveImageDigests() *Config {
	c.Command += helpers.Option(resolveImageDigests)
	return c
}

// Services - Print the service names, one per line
func (c *Config) Services() *Config {
	c.Command += helpers.Option(services)
	return c
}

// Variables - Print model variables and default values
func (c *Config) Variables() *Config {
	c.Command += helpers.Option(variables)
	return c
}

// Volumes - Print the volume names, one per line
func (c *Config) Volumes() *Config {
	c.Command += helpers.Option(volumes)
	return c
}

// GetCommand - Get the complete command string
func (c *Config) GetCommand() string {
	return c.Command
}

// Exec - Execute the command
func (c *Config) Exec() {
	helpers.GeneralExec(c.Command, false)
}

// ExecWithPrivileges - Execute the command with elevated privileges
func (c *Config) ExecWithPrivileges() {
	helpers.GeneralExec(c.Command, true)
}
