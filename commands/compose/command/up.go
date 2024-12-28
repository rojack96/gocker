package command

import (
	"github.com/rojack96/gocker/commands/common"
	"github.com/rojack96/gocker/helpers"
	"strconv"
)

const (
	abortOnContainerExit    = "--abort-on-container-exit"
	abortOnContainerFailure = "--abort-on-container-failure"
	alwaysRecreateDeps      = "--always-recreate-deps"
	attach                  = "--attach"
	attachDependencies      = "--attach-dependencies"
	exitCodeFrom            = "--exit-code-from"
	noAttach                = "--no-attach"
	noStart                 = "--no-start"
	renewAnonVolumes        = "--renew-anon-volumes"
	wait                    = "--wait"
	waitTimeout             = "--wait-timeout"
	watch                   = "--watch"
)

type Up struct {
	command string
}

func NewUp(cmd string) *Up {
	return &Up{command: cmd}
}

// AbortOnContainerExit - Stops all containers if any container was stopped. Incompatible with -d
// TODO Implement the possibility to make it incompatible with Detach
func (up *Up) AbortOnContainerExit() *Up {
	return &Up{command: up.command + helpers.Option(abortOnContainerExit)}
}

// AbortOnContainerFailure - Stops all containers if any container exited with failure. Incompatible with -d
func (up *Up) AbortOnContainerFailure() *Up {
	return &Up{command: up.command + helpers.Option(abortOnContainerFailure)}
}

// AlwaysRecreateDeps - Recreate dependent containers. Incompatible with --no-recreate.
func (up *Up) AlwaysRecreateDeps() *Up {
	return &Up{command: up.command + helpers.Option(alwaysRecreateDeps)}
}

// Attach - Restrict attaching to the specified services. Incompatible with --attach-dependencies.
func (up *Up) Attach(services []string) *Up {
	return &Up{command: up.command + helpers.StringArray(attach, services...)}
}

// AttachDependencies - Automatically attach to log output of dependent services
func (up *Up) AttachDependencies() *Up {
	return &Up{command: up.command + helpers.StringArray(attachDependencies)}
}

// Build - Build images before starting containers
func (up *Up) Build() *Up {
	return &Up{command: up.command + common.Build()}
}

// Detach - Detached mode: Run containers in the background
func (up *Up) Detach() *Up {
	return &Up{command: up.command + common.Detach()}
}

// DryRun - Execute command in dry run mode
func (up *Up) DryRun() *Up {
	return &Up{command: up.command + common.DryRun()}
}

// ExitCodeFrom - Return the exit code of the selected service container. Implies --abort-on-container-exit
func (up *Up) ExitCodeFrom(service string) *Up {
	return &Up{command: up.command + helpers.String(exitCodeFrom, service)}
}

// ForceRecreate - Recreate containers even if their configuration and image haven't changed
func (up *Up) ForceRecreate() *Up {
	return &Up{command: up.command + common.ForceRecreate()}
}

// NoAttach - Do not attach (stream logs) to the specified services
func (up *Up) NoAttach(services []string) *Up {
	return &Up{command: up.command + helpers.StringArray(noAttach, services...)}
}

// NoBuild - Don't build an image, even if it's policy
func (up *Up) NoBuild() *Up {
	return &Up{command: up.command + common.NoBuild()}
}

// NoColor - Produce monochrome output
func (up *Up) NoColor() *Up {
	return &Up{command: up.command + common.NoColor()}
}

// NoDeps - Don't start linked services
func (up *Up) NoDeps() *Up {
	return &Up{command: up.command + common.NoDeps()}
}

// NoLogPrefix - Don't print prefix in logs
func (up *Up) NoLogPrefix() *Up {
	return &Up{command: up.command + common.NoLogPrefix()}
}

// NoRecreate -  If containers already exist, don't recreate them. Incompatible with --force-recreate.
func (up *Up) NoRecreate() *Up {
	return &Up{command: up.command + common.NoRecreate()}
}

// NoStart - Don't start the services after creating them
func (up *Up) NoStart() *Up {
	return &Up{command: up.command + helpers.Option(noStart)}
}

// Pull - Pull image before running ("always"|"missing"|"never") (default "policy")
func (up *Up) Pull(pullPolicy string) *Up {
	return &Up{command: up.command + helpers.String(pull, pullPolicy)}
}

// QuietPull - Pull without printing progress information
func (up *Up) QuietPull() *Up {
	return &Up{command: up.command + common.QuietPull()}
}

// RemoveOrphans - Remove containers for services not defined in the Compose file
func (up *Up) RemoveOrphans() *Up {
	return &Up{command: up.command + common.RemoveOrphans()}
}

// RenewAnonVolumes - Recreate anonymous volumes instead of retrieving data from the previous containers
func (up *Up) RenewAnonVolumes() *Up {
	return &Up{command: up.command + helpers.Option(renewAnonVolumes)}

}

// Scale - Scale SERVICE to NUM instances. Overrides the scale setting in the Compose file if present.
func (up *Up) Scale(service string, instances int) *Up {
	serviceScale := service + "=" + strconv.Itoa(instances)
	return &Up{command: up.command + serviceScale}
}

// Timeout - Use this timeout in seconds for container shutdown when attached or when containers are already running
func (up *Up) Timeout(seconds int) *Up {
	return &Up{command: up.command + common.Timeout(seconds)}
}

// Timestamps - Show timestamps
func (up *Up) Timestamps() *Up {
	return &Up{command: up.command + common.Timestamps()}
}

// Wait - Wait for services to be running|healthy. Implies detached mode.
func (up *Up) Wait() *Up {
	return &Up{command: up.command + helpers.Option(wait)}
}

// WaitTimeout - Maximum duration to wait for the project to be running|healthy
func (up *Up) WaitTimeout(seconds int) *Up {
	return &Up{command: up.command + helpers.Int(waitTimeout, seconds)}
}

// Watch - Watch source code and rebuild/refresh containers when files are updated.
func (up *Up) Watch() *Up {
	return &Up{command: up.command + helpers.Option(watch)}
}

func (up *Up) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return common.SetCommand(up.command + helpers.ServiceName(serviceNames...))
}
