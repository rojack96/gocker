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
	return &Config{Command: c.Command + option.DryRun()}
}

// Format - Format the output. Values: [yaml | json] (default "yaml")
func (c *Config) Format(value string) *Config {
	return &Config{Command: c.Command + common.Format(value)}
}

// Hash - Print the service config hash, one per line
func (c *Config) Hash(value string) *Config {
	return &Config{Command: c.Command + helpers.String(hash, value)}
}

// Images - Print the image names, one per line
func (c *Config) Images() *Config {
	return &Config{Command: c.Command + helpers.Option(images)}
}

// NoConsistency - Don't check model consistency
func (c *Config) NoConsistency() *Config {
	return &Config{Command: c.Command + helpers.Option(noConsistency)}
}

// NoInterpolate - Don't interpolate environment variables
func (c *Config) NoInterpolate() *Config {
	return &Config{Command: c.Command + helpers.Option(noInterpolate)}
}

// NoNormalize - Don't normalize compose model
func (c *Config) NoNormalize() *Config {
	return &Config{Command: c.Command + helpers.Option(noNormalize)}
}

// NoPathResolution - Don't resolve file paths
func (c *Config) NoPathResolution() *Config {
	return &Config{Command: c.Command + helpers.Option(noPathResolution)}
}

// Output - Save to file (default to stdout)
func (c *Config) Output(value string) *Config {
	return &Config{Command: c.Command + helpers.String(output, value)}
}

// Profiles - Print the profile names, one per line
func (c *Config) Profiles() *Config {
	return &Config{Command: c.Command + helpers.Option(profiles)}
}

// Quiet - Only validate the configuration, don't print anything
func (c *Config) Quiet() *Config {
	return &Config{Command: c.Command + helpers.Option(quiet)}
}

// ResolveImageDigests - Pin image tags to digests
func (c *Config) ResolveImageDigests() *Config {
	return &Config{Command: c.Command + helpers.Option(resolveImageDigests)}

}

// Services - Print the service names, one per line
func (c *Config) Services() *Config {
	return &Config{Command: c.Command + helpers.Option(services)}

}

// Variables - Print model variables and default values
func (c *Config) Variables() *Config {
	return &Config{Command: c.Command + helpers.Option(variables)}

}

// Volumes - Print the volume names, one per line
func (c *Config) Volumes() *Config {
	return &Config{Command: c.Command + helpers.Option(volumes)}

}

func (c *Config) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: c.Command + helpers.ServiceName(serviceNames...)}
}
