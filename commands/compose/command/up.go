package command

import (
	"fmt"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
)

const (
	abortOnContainerExit = "--abort-on-container-exit"
	/*Options:
	      --abort-on-container-exit
	      --abort-on-container-failure   Stops all containers if any container exited with failure. Incompatible with -d
	      --always-recreate-deps         Recreate dependent containers. Incompatible with --no-recreate.
	      --attach stringArray           Restrict attaching to the specified services. Incompatible with --attach-dependencies.
	      --attach-dependencies          Automatically attach to log output of dependent services
	      --build                        Build images before starting containers
	  -d, --detach                       Detached mode: Run containers in the background
	      --dry-run                      Execute command in dry run mode
	      --exit-code-from string        Return the exit code of the selected service container. Implies --abort-on-container-exit
	      --force-recreate               Recreate containers even if their configuration and image haven't changed
	      --no-attach stringArray        Do not attach (stream logs) to the specified services
	      --no-build                     Don't build an image, even if it's policy
	      --no-color                     Produce monochrome output
	      --no-deps                      Don't start linked services
	      --no-log-prefix                Don't print prefix in logs
	      --no-recreate                  If containers already exist, don't recreate them. Incompatible with --force-recreate.
	      --no-start                     Don't start the services after creating them
	      --pull string                  Pull image before running ("always"|"missing"|"never") (default "policy")
	      --quiet-pull                   Pull without printing progress information
	      --remove-orphans               Remove containers for services not defined in the Compose file
	  -V, --renew-anon-volumes           Recreate anonymous volumes instead of retrieving data from the previous containers
	      --scale scale                  Scale SERVICE to NUM instances. Overrides the scale setting in the Compose file if present.
	  -t, --timeout int                  Use this timeout in seconds for container shutdown when attached or when containers are already running
	      --timestamps                   Show timestamps
	      --wait                         Wait for services to be running|healthy. Implies detached mode.
	      --wait-timeout int             Maximum duration to wait for the project to be running|healthy
	  -w, --watch                        Watch source code and rebuild/refresh containers when files are updated.*/
)

type Up struct {
	Command string
}

// AbortOnContainerExit Stops all containers if any container was stopped. Incompatible with -d
// TODO Implement the possibility to make it incompatible with Detach
func (up *Up) AbortOnContainerExit() *Up {
	up.Command += helpers.Option(abortOnContainerExit)
	return up
}

// TODO continue from here

func (up *Up) AbortOnContainerFailure() *Up {
	up.Command += " --abort-on-container-failure"
	return up
}

func (up *Up) AlwaysRecreateDeps() *Up {
	up.Command += " --always-recreate-deps"
	return up
}

func (up *Up) Attach(services []string) *Up {
	for _, service := range services {
		up.Command += " --attach " + service
	}
	return up
}

func (up *Up) AttachDependencies() *Up {
	up.Command += " --attach-dependencies"
	return up
}

func (up *Up) Build() *Up {
	up.Command += " --command"
	return up
}

func (up *Up) Detach() *Up {
	up.Command += " -d"
	return up
}

func (up *Up) DryRun() *Up {
	up.Command += option.DryRun()
	return up
}

func (up *Up) ExitCodeFrom(service string) *Up {
	up.Command += " --exit-code-from " + service
	return up
}

func (up *Up) ForceRecreate() *Up {
	up.Command += " --force-recreate"
	return up
}

func (up *Up) NoAttach(services []string) *Up {
	for _, service := range services {
		up.Command += " --no-attach " + service
	}
	return up
}

func (up *Up) NoBuild() *Up {
	up.Command += " --no-command"
	return up
}

func (up *Up) NoColor() *Up {
	up.Command += " --no-color"
	return up
}

func (up *Up) NoDeps() *Up {
	up.Command += " --no-deps"
	return up
}

func (up *Up) NoLogPrefix() *Up {
	up.Command += " --no-log-prefix"
	return up
}

func (up *Up) NoRecreate() *Up {
	up.Command += " --no-recreate"
	return up
}

func (up *Up) NoStart() *Up {
	up.Command += " --no-start"
	return up
}

func (up *Up) Pull(pullPolicy string) *Up {
	up.Command += " --pull " + pullPolicy
	return up
}

func (up *Up) QuietPull() *Up {
	up.Command += " --quiet-pull"
	return up
}

func (up *Up) RemoveOrphans() *Up {
	up.Command += " --remove-orphans"
	return up
}

func (up *Up) RenewAnonVolumes() *Up {
	up.Command += " -V"
	return up
}

func (up *Up) Scale(service string, instances int) *Up {
	up.Command += " --scale " + fmt.Sprintf("%s=%d", service, instances)
	return up
}

func (up *Up) Timeout(seconds int) *Up {
	up.Command += " -t " + fmt.Sprintf("%d", seconds)
	return up
}

func (up *Up) Timestamps() *Up {
	up.Command += " --timestamps"
	return up
}

func (up *Up) Wait() *Up {
	up.Command += " --wait"
	return up
}

func (up *Up) WaitTimeout(seconds int) *Up {
	up.Command += " --wait-timeout " + fmt.Sprintf("%d", seconds)
	return up

}

func (up *Up) Watch() *Up {
	up.Command += " -w"
	return up
}

func (up *Up) ServiceName(serviceName ...string) *Up {
	up.Command += helpers.ServiceName(serviceName...)
	return up
}

func (up *Up) GetCommand() string {
	return up.Command
}

func (up *Up) Exec() {
	helpers.GeneralExec(up.Command)
}
