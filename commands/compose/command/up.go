package command

import (
	"github.com/rojack96/gocker/commands/compose/common"
	"github.com/rojack96/gocker/commands/compose/option"
	"github.com/rojack96/gocker/helpers"
	"strconv"
)

const (
	abortOnContainerExit    = "--abort-on-container-exit"
	abortOnContainerFailure = "--abort-on-container-failure"
	alwaysRecreateDeps      = "--always-recreate-deps"
	attach                  = "--attach"
	attachDependencies      = "--attach-dependencies"
	detach                  = "--detach"
	exitCodeFrom            = "--exit-code-from"
	noAttach                = "--no-attach"
	noDeps                  = "--no-deps"
	noStart                 = "--no-start"
	renewAnonVolumes        = "--renew-anon-volumes"
	timeout                 = "--timeout"
	wait                    = "--wait"
	waitTimeout             = "--wait-timeout"
	watch                   = "--watch"
)

type Up struct {
	Command string
}

// AbortOnContainerExit - Stops all containers if any container was stopped. Incompatible with -d
// TODO Implement the possibility to make it incompatible with Detach
func (up *Up) AbortOnContainerExit() *Up {
	up.Command += helpers.Option(abortOnContainerExit)
	return up
}

// AbortOnContainerFailure - Stops all containers if any container exited with failure. Incompatible with -d
func (up *Up) AbortOnContainerFailure() *Up {
	up.Command += helpers.Option(abortOnContainerFailure)
	return up
}

// AlwaysRecreateDeps - Recreate dependent containers. Incompatible with --no-recreate.
func (up *Up) AlwaysRecreateDeps() *Up {
	up.Command += helpers.Option(alwaysRecreateDeps)
	return up
}

// Attach - Restrict attaching to the specified services. Incompatible with --attach-dependencies.
func (up *Up) Attach(services []string) *Up {
	up.Command += helpers.StringArray(attach, services...)
	return up
}

// AttachDependencies - Automatically attach to log output of dependent services
func (up *Up) AttachDependencies() *Up {
	up.Command += helpers.StringArray(attachDependencies)
	return up
}

// Build - Build images before starting containers
func (up *Up) Build() *Up {
	up.Command += option.Build()
	return up
}

// Detach - Detached mode: Run containers in the background
func (up *Up) Detach() *Up {
	up.Command += helpers.Option(detach)
	return up
}

// DryRun - Execute command in dry run mode
func (up *Up) DryRun() *Up {
	up.Command += option.DryRun()
	return up
}

// ExitCodeFrom - Return the exit code of the selected service container. Implies --abort-on-container-exit
func (up *Up) ExitCodeFrom(service string) *Up {
	up.Command += helpers.String(exitCodeFrom, service)
	return up
}

// ForceRecreate - Recreate containers even if their configuration and image haven't changed
func (up *Up) ForceRecreate() *Up {
	up.Command += option.ForceRecreate()
	return up
}

// NoAttach - Do not attach (stream logs) to the specified services
func (up *Up) NoAttach(services []string) *Up {
	up.Command += helpers.StringArray(noAttach, services...)
	return up
}

// NoBuild - Don't build an image, even if it's policy
func (up *Up) NoBuild() *Up {
	up.Command += option.NoBuild()
	return up
}

// NoColor - Produce monochrome output
func (up *Up) NoColor() *Up {
	up.Command += option.NoColor()
	return up
}

// NoDeps - Don't start linked services
func (up *Up) NoDeps() *Up {
	up.Command += helpers.Option(noDeps)
	return up
}

// NoLogPrefix - Don't print prefix in logs
func (up *Up) NoLogPrefix() *Up {
	up.Command += option.NoLogPrefix()
	return up
}

// NoRecreate -  If containers already exist, don't recreate them. Incompatible with --force-recreate.
func (up *Up) NoRecreate() *Up {
	up.Command += option.NoRecreate()
	return up
}

// NoStart - Don't start the services after creating them
func (up *Up) NoStart() *Up {
	up.Command += helpers.Option(noStart)
	return up
}

// Pull - Pull image before running ("always"|"missing"|"never") (default "policy")
func (up *Up) Pull(pullPolicy string) *Up {
	up.Command += helpers.String(pull, pullPolicy)
	return up
}

// QuietPull - Pull without printing progress information
func (up *Up) QuietPull() *Up {
	return &Up{Command: up.Command + option.QuietPull()}
}

// RemoveOrphans - Remove containers for services not defined in the Compose file
func (up *Up) RemoveOrphans() *Up {
	return &Up{Command: up.Command + option.RemoveOrphans()}
}

// RenewAnonVolumes - Recreate anonymous volumes instead of retrieving data from the previous containers
func (up *Up) RenewAnonVolumes() *Up {
	up.Command += helpers.Option(renewAnonVolumes)
	return up
}

// Scale - Scale SERVICE to NUM instances. Overrides the scale setting in the Compose file if present.
func (up *Up) Scale(service string, instances int) *Up {
	serviceScale := service + "=" + strconv.Itoa(instances)
	return &Up{Command: up.Command + serviceScale}
}

// Timeout - Use this timeout in seconds for container shutdown when attached or when containers are already running
func (up *Up) Timeout(seconds int) *Up {
	up.Command += helpers.Int(timeout, seconds)
	return up
}

// Timestamps - Show timestamps
func (up *Up) Timestamps() *Up {
	up.Command += option.Timestamps()
	return up
}

// Wait - Wait for services to be running|healthy. Implies detached mode.
func (up *Up) Wait() *Up {
	up.Command += helpers.Option(wait)
	return up
}

// WaitTimeout - Maximum duration to wait for the project to be running|healthy
func (up *Up) WaitTimeout(seconds int) *Up {
	up.Command += helpers.Int(waitTimeout, seconds)
	return up

}

// Watch - Watch source code and rebuild/refresh containers when files are updated.
func (up *Up) Watch() *Up {
	up.Command += helpers.Option(watch)
	return up
}

func (up *Up) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	up.Command += helpers.ServiceName(serviceNames...)
	return &common.CommandExecutor{Command: up.Command}
}
