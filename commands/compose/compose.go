package compose

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/commands/compose/command"
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
)

type Compose struct {
	command string
}

func New(cmd string) *Compose {
	return &Compose{command: cmd}
}

/*
	------------------------------
	|          OPTIONS           |
	------------------------------
*/

// AllResources - Include all resources, even those not used by services
func (c *Compose) AllResources() *Compose {
	return &Compose{command: c.command + helpers.Option(allResources)}
}

// Ansi - Control when to print ANSI control characters ("never"|"always"|"auto") (default "auto")
func (c *Compose) Ansi(ansiString string) *Compose {
	return &Compose{command: c.command + helpers.String(ansi, ansiString)}

}

// Compatibility - Run compose in backward compatibility mode
func (c *Compose) Compatibility() *Compose {
	return &Compose{command: c.command + helpers.Option(compatibility)}
}

// DryRun - Execute command in dry run mode
func (c *Compose) DryRun() *Compose {
	return &Compose{command: c.command + common.DryRun()}
}

// EnvFile - Specify an alternate environment file
func (c *Compose) EnvFile(files ...string) *Compose {
	return &Compose{command: c.command + helpers.StringArray(envFile, files...)}
}

// FileName Compose - configuration files
func (c *Compose) FileName(files ...string) *Compose {
	return &Compose{command: c.command + helpers.StringArray(file, files...)}
}

// Parallelism - Control max parallelism, -1 for unlimited (default -1)
func (c *Compose) Parallelism(parallelismValue int) *Compose {
	return &Compose{command: c.command + helpers.Int(parallelism, parallelismValue)}
}

// Profile - Specify a profile to enable
func (c *Compose) Profile(profiles ...string) *Compose {
	return &Compose{command: c.command + helpers.StringArray(profile, profiles...)}
}

// Progress - Set type of progress output (auto, tty, plain, quiet) (default "auto")
func (c *Compose) Progress(prog string) *Compose {
	return &Compose{command: c.command + helpers.String(progress, prog)}
}

// ProjectDirectory - Specify an alternate working directory
// (default: the path of the first specified, Compose file)
func (c *Compose) ProjectDirectory(pd string) *Compose {
	return &Compose{command: c.command + helpers.String(projectDirectory, pd)}
}

// ProjectName - Project name
func (c *Compose) ProjectName(pn string) *Compose {
	return &Compose{command: c.command + helpers.String(projectName, pn)}
}

/*
	------------------------------
	|          COMMANDS          |
	------------------------------
*/

// Attach - Attach local standard input, output, and error streams to a service's running container
func (c *Compose) Attach() *command.Attach {
	return command.NewAttach(c.command + helpers.Command(attach))
}

// Build - Build or rebuild services
func (c *Compose) Build() *command.Build {
	return command.NewBuild(c.command + helpers.Command(build))
}

// Config - Parse, resolve and render compose file in canonical format
func (c *Compose) Config() *command.Config {
	return command.NewConfig(c.command + helpers.Command(config))
}

// Cp - Copy files/folders between a service container and the local filesystem
func (c *Compose) Cp() *command.Cp { return command.NewCp(c.command + helpers.Command(cp)) }

// Create - Creates containers for a service
func (c *Compose) Create() *command.Create {
	return command.NewCreate(c.command + helpers.Command(create))
}

// Down - Stop and remove containers, networks
func (c *Compose) Down() *command.Down {
	return command.NewDown(c.command + helpers.Command(down))
}

// Events - Receive real time events from containers
func (c *Compose) Events() *command.Events {
	return command.NewEvents(c.command + helpers.Command(events))
}

// Exec - Receive real time events from containers
func (c *Compose) Exec() *command.Exec {
	return command.NewExec(c.command + helpers.Command(exec))
}

// Images - List images used by the created containers
func (c *Compose) Images() *command.Images {
	return command.NewImages(c.command + helpers.Command(images))
}

// Kill - Force stop service containers
func (c *Compose) Kill() *command.Kill {
	return command.NewKill(c.command + helpers.Command(kill))
}

// Logs - View output from containers
func (c *Compose) Logs() *command.Logs {
	return command.NewLogs(c.command + helpers.Command(logs))
}

// Ls - List running compose projects
func (c *Compose) Ls() *command.Ls {
	return command.NewLs(c.command + helpers.Command(ls))
}

// Pause - Pause services
func (c *Compose) Pause() *command.Pause {
	return command.NewPause(c.command + helpers.Command(pause))
}

// Port - Print the public port for a port binding
func (c *Compose) Port() *command.Port {
	return command.NewPort(c.command + helpers.Command(port))
}

// Ps - List containers
func (c *Compose) Ps() *command.Ps {
	return command.NewPs(c.command + helpers.Command(ps))
}

// Pull - Pull service images
func (c *Compose) Pull() *command.Pull {
	return command.NewPull(c.command + helpers.Command(pull))
}

// Push - Push service images
func (c *Compose) Push() *command.Push {
	return command.NewPush(c.command + helpers.Command(push))
}

// Restart - Restart service containers
func (c *Compose) Restart() *command.Restart {
	return command.NewRestart(c.command + helpers.Command(restart))
}

// Rm - Removes stopped service containers
func (c *Compose) Rm() *command.Rm {
	return command.NewRm(c.command + helpers.Command(rm))
}

// Run - Run a one-off command on a service
func (c *Compose) Run() *command.Run {
	return command.NewRun(c.command + helpers.Command(run))
}

// Scale - Scale services
func (c *Compose) Scale() *command.Scale {
	return command.NewScale(c.command + helpers.Command(scale))
}

// Start - Start services
func (c *Compose) Start() *command.Start {
	return command.NewStart(c.command + helpers.Command(start))
}

// Stats - Display a live stream of container(s) resource usage statistics
func (c *Compose) Stats() *command.Stats {
	return command.NewStats(c.command + helpers.Command(stats))
}

// Stop - Stop services
func (c *Compose) Stop() *command.Stop {
	return command.NewStop(c.command + helpers.Command(stop))
}

// Top - Display the running processes
func (c *Compose) Top() *command.Top {
	return command.NewTop(c.command + helpers.Command(top))
}

func (c *Compose) Unpause() *command.Unpause {
	return command.NewUnpause(c.command + helpers.Command(unpause))
}

// Up - Create and start containers
func (c *Compose) Up() *command.Up { return command.NewUp(c.command + helpers.Command(up)) }

// Version - Show the Docker Compose version information
func (c *Compose) Version() *command.Version {
	return command.NewVersion(c.command + helpers.Command(version))
}

// Wait - Block until the first service container stops
func (c *Compose) Wait() *command.Wait {
	return command.NewWait(c.command + helpers.Command(wait))
}

// Watch - Watch command context for service and rebuild/refresh containers when files are updated
func (c *Compose) Watch() *command.Watch {
	return command.NewWatch(c.command + helpers.Command(watch))
}
