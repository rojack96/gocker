package compose

import (
	"github.com/rojack96/gocker/commands/compose/command"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	// Options:
	allResources     = "--all-resources"
	ansi             = "--ansi"
	compatibility    = "--compatibility"
	envFile          = "--env-file"
	file             = "--file"
	parallelism      = "--parallel"
	profile          = "--profile"
	progress         = "--progress"
	projectDirectory = "--project-directory"
	projectName      = "--project-name"

	// Commands:
	attach  = "attach"
	build   = "build"
	config  = "config"
	cp      = "cp"
	create  = "create"
	down    = "down"
	events  = "events"
	exec    = "exec"
	images  = "images"
	kill    = "kill"
	logs    = "logs"
	ls      = "ls"
	pause   = "pause"
	port    = "port"
	ps      = "ps"
	pull    = "pull"
	push    = "push"
	restart = "restart"
	rm      = "rm"
	run     = "run"
	scale   = "scale"
	start   = "start"
	stats   = "stats"
	stop    = "stop"
	top     = "top"
	unpause = "unpause"
	up      = "up"
	version = "version"
	wait    = "wait"
	watch   = "watch"
	/*
		Commands:

		port        Print the public port for a port binding
		ps          List containers
		pull        Pull service images
		push        Push service images
		restart     Restart service containers
		rm          Removes stopped service containers
		run         Run a one-off command on a service
		scale       Scale services
		start       Start services
		stats       Display a live stream of container(s) resource usage statistics
		stop        Stop services
		top         Display the running processes
		unpause     Unpause services
		version     Show the Docker Compose version information
		wait        Block until the first service container stops
		watch       Watch command context for service and rebuild/refresh containers when files are updated
	*/
)

type Compose struct {
	Command string
}

/*
	------------------------------
	|          OPTIONS           |
	------------------------------
*/

// AllResources - Include all resources, even those not used by services
func (c *Compose) AllResources() *Compose {
	c.Command += helpers.Option(allResources)
	return c
}

// Ansi - Control when to print ANSI control characters ("never"|"always"|"auto") (default "auto")
func (c *Compose) Ansi(ansiString string) *Compose {
	c.Command += helpers.String(ansi, ansiString)
	return c
}

// Compatibility - Run compose in backward compatibility mode
func (c *Compose) Compatibility() *Compose {
	c.Command += helpers.Option(compatibility)
	return c
}

// DryRun - Execute command in dry run mode
func (c *Compose) DryRun() *Compose {
	c.Command += option.DryRun()
	return c
}

// EnvFile - Specify an alternate environment file
func (c *Compose) EnvFile(files ...string) *Compose {
	c.Command += helpers.StringArray(envFile, files...)
	return c
}

// FileName Compose - configuration files
func (c *Compose) FileName(files ...string) *Compose {
	c.Command += helpers.StringArray(file, files...)
	return c
}

// Parallelism - Control max parallelism, -1 for unlimited (default -1)
func (c *Compose) Parallelism(parallelismValue int) *Compose {
	c.Command += helpers.Int(parallelism, parallelismValue)
	return c
}

// Profile - Specify a profile to enable
func (c *Compose) Profile(profiles ...string) *Compose {
	c.Command += helpers.StringArray(profile, profiles...)
	return c
}

// Progress - Set type of progress output (auto, tty, plain, quiet) (default "auto")
func (c *Compose) Progress(prog string) *Compose {
	c.Command += helpers.String(progress, prog)
	return c
}

// ProjectDirectory - Specify an alternate working directory
// (default: the path of the first specified, Compose file)
func (c *Compose) ProjectDirectory(pd string) *Compose {
	c.Command += helpers.String(projectDirectory, pd)
	return c
}

// ProjectName - Project name
func (c *Compose) ProjectName(pn string) *Compose {
	c.Command += helpers.String(projectName, pn)
	return c
}

/*
	------------------------------
	|          COMMANDS          |
	------------------------------
*/

// Attach - Attach local standard input, output, and error streams to a service's running container
func (c *Compose) Attach() *command.Attach {
	return &command.Attach{Command: c.Command + helpers.Command(attach)}
}

// Build - Build or rebuild services
func (c *Compose) Build() *command.Build {
	return &command.Build{Command: c.Command + helpers.Command(build)}
}

// Config - Parse, resolve and render compose file in canonical format
func (c *Compose) Config() *command.Config {
	return &command.Config{Command: c.Command + helpers.Command(config)}
}

// Cp - Copy files/folders between a service container and the local filesystem
func (c *Compose) Cp() *command.Cp { return &command.Cp{Command: c.Command + helpers.Command(cp)} }

// Create - Creates containers for a service
func (c *Compose) Create() *command.Create {
	return &command.Create{Command: c.Command + helpers.Command(create)}
}

// Down - Stop and remove containers, networks
func (c *Compose) Down() *command.Down {
	return &command.Down{Command: c.Command + helpers.Command(down)}
}

// Events - Receive real time events from containers
func (c *Compose) Events() *command.Events {
	return &command.Events{Command: c.Command + helpers.Command(events)}
}

// Exec - Receive real time events from containers
func (c *Compose) Exec() *command.Exec {
	return &command.Exec{Command: c.Command + helpers.Command(exec)}
}

// Images - List images used by the created containers
func (c *Compose) Images() *command.Images {
	return &command.Images{Command: c.Command + helpers.Command(images)}
}

// Kill - Force stop service containers
func (c *Compose) Kill() *command.Kill {
	return &command.Kill{Command: c.Command + helpers.Command(kill)}
}

// Ls - List running compose projects
func (c *Compose) Ls() *command.Ls {
	return &command.Ls{Command: c.Command + helpers.Command(ls)}
}

// Pause - Pause services
func (c *Compose) Pause() *command.Pause {
	return &command.Pause{Command: c.Command + helpers.Command(pause)}
}

// Up - Create and start containers
func (c *Compose) Up() *command.Up { return &command.Up{Command: c.Command + helpers.Command(up)} }

// Logs - View output from containers
func (c *Compose) Logs() *command.Logs {
	return &command.Logs{Command: c.Command + helpers.Command(logs)}
}
