package gocker

import (
	"gocker/commands"
	"gocker/helpers"
)

type Compose struct {
	Command string
}

// Options

func (c *Compose) FileName(fileName string) *Compose {
	c.Command += " -f" + helpers.AppendString(fileName)

	return c
}

func (c *Compose) EnvFile(fileName string) *Compose {
	c.Command += " --env-file" + helpers.AppendString(fileName)
	return c
}

// TODO implement all options

/* Commands */

// Build or rebuild services
func (c *Compose) Build() *commands.Build {
	return &commands.Build{Command: c.Command + " build"}
}

// Up Create and start containers
func (c *Compose) Up() *commands.Up {
	return &commands.Up{Command: c.Command + " up"}
}

// TODO implement all Commands

/*
Options:
--all-resources              Include all resources, even those not used by services
--ansi string                Control when to print ANSI control characters
("never"|"always"|"auto") (default "auto")
--compatibility              Run compose in backward compatibility mode
--dry-run                    Execute command in dry run mode
--env-file stringArray       Specify an alternate environment file
-f, --file stringArray           Compose configuration files
--parallel int               Control max parallelism, -1 for unlimited (default -1)
--profile stringArray        Specify a profile to enable
--progress string            Set type of progress output (auto, tty, plain, quiet) (default "auto")
--project-directory string   Specify an alternate working directory
(default: the path of the, first specified, Compose file)
-p, --project-name string        Project name

Commands:
attach      Attach local standard input, output, and error streams to a service's running container
build       Build or rebuild services
config      Parse, resolve and render compose file in canonical format
cp          Copy files/folders between a service container and the local filesystem
create      Creates containers for a service
down        Stop and remove containers, networks
events      Receive real time events from containers
exec        Execute a command in a running container
images      List images used by the created containers
kill        Force stop service containers
logs        View output from containers
ls          List running compose projects
pause       Pause services
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
up          Create and start containers
version     Show the Docker Compose version information
wait        Block until the first service container stops
watch       Watch build context for service and rebuild/refresh containers when files are updated
*/
