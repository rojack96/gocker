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
	exitCodeFrom            = "--exit-code-from"
	noAttach                = "--no-attach"
	noDeps                  = "--no-deps"
	noStart                 = "--no-start"
	renewAnonVolumes        = "--renew-anon-volumes"
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
	return &Up{Command: up.Command + helpers.Option(abortOnContainerExit)}
}

// AbortOnContainerFailure - Stops all containers if any container exited with failure. Incompatible with -d
func (up *Up) AbortOnContainerFailure() *Up {
	return &Up{Command: up.Command + helpers.Option(abortOnContainerFailure)}
}

// AlwaysRecreateDeps - Recreate dependent containers. Incompatible with --no-recreate.
func (up *Up) AlwaysRecreateDeps() *Up {
	return &Up{Command: up.Command + helpers.Option(alwaysRecreateDeps)}
}

// Attach - Restrict attaching to the specified services. Incompatible with --attach-dependencies.
func (up *Up) Attach(services []string) *Up {
	return &Up{Command: up.Command + helpers.StringArray(attach, services...)}
}

// AttachDependencies - Automatically attach to log output of dependent services
func (up *Up) AttachDependencies() *Up {
	return &Up{Command: up.Command + helpers.StringArray(attachDependencies)}
}

// Build - Build images before starting containers
func (up *Up) Build() *Up {
	return &Up{Command: up.Command + option.Build()}
}

// Detach - Detached mode: Run containers in the background
func (up *Up) Detach() *Up {
	return &Up{Command: up.Command + option.Detach()}
}

// DryRun - Execute command in dry run mode
func (up *Up) DryRun() *Up {
	return &Up{Command: up.Command + option.DryRun()}
}

// ExitCodeFrom - Return the exit code of the selected service container. Implies --abort-on-container-exit
func (up *Up) ExitCodeFrom(service string) *Up {
	return &Up{Command: up.Command + helpers.String(exitCodeFrom, service)}
}

// ForceRecreate - Recreate containers even if their configuration and image haven't changed
func (up *Up) ForceRecreate() *Up {
	return &Up{Command: up.Command + option.ForceRecreate()}
}

// NoAttach - Do not attach (stream logs) to the specified services
func (up *Up) NoAttach(services []string) *Up {
	return &Up{Command: up.Command + helpers.StringArray(noAttach, services...)}
}

// NoBuild - Don't build an image, even if it's policy
func (up *Up) NoBuild() *Up {
	return &Up{Command: up.Command + option.NoBuild()}
}

// NoColor - Produce monochrome output
func (up *Up) NoColor() *Up {
	return &Up{Command: up.Command + option.NoColor()}
}

// NoDeps - Don't start linked services
func (up *Up) NoDeps() *Up {
	return &Up{Command: up.Command + helpers.Option(noDeps)}
}

// NoLogPrefix - Don't print prefix in logs
func (up *Up) NoLogPrefix() *Up {
	return &Up{Command: up.Command + option.NoLogPrefix()}
}

// NoRecreate -  If containers already exist, don't recreate them. Incompatible with --force-recreate.
func (up *Up) NoRecreate() *Up {
	return &Up{Command: up.Command + option.NoRecreate()}
}

// NoStart - Don't start the services after creating them
func (up *Up) NoStart() *Up {
	return &Up{Command: up.Command + helpers.Option(noStart)}
}

// Pull - Pull image before running ("always"|"missing"|"never") (default "policy")
func (up *Up) Pull(pullPolicy string) *Up {
	return &Up{Command: up.Command + helpers.String(pull, pullPolicy)}
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
	return &Up{Command: up.Command + helpers.Option(renewAnonVolumes)}

}

// Scale - Scale SERVICE to NUM instances. Overrides the scale setting in the Compose file if present.
func (up *Up) Scale(service string, instances int) *Up {
	serviceScale := service + "=" + strconv.Itoa(instances)
	return &Up{Command: up.Command + serviceScale}
}

// Timeout - Use this timeout in seconds for container shutdown when attached or when containers are already running
func (up *Up) Timeout(seconds int) *Up {
	return &Up{Command: up.Command + common.Timeout(seconds)}
}

// Timestamps - Show timestamps
func (up *Up) Timestamps() *Up {
	return &Up{Command: up.Command + option.Timestamps()}
}

// Wait - Wait for services to be running|healthy. Implies detached mode.
func (up *Up) Wait() *Up {
	return &Up{Command: up.Command + helpers.Option(wait)}
}

// WaitTimeout - Maximum duration to wait for the project to be running|healthy
func (up *Up) WaitTimeout(seconds int) *Up {
	return &Up{Command: up.Command + helpers.Int(waitTimeout, seconds)}
}

// Watch - Watch source code and rebuild/refresh containers when files are updated.
func (up *Up) Watch() *Up {
	return &Up{Command: up.Command + helpers.Option(watch)}
}

func (up *Up) ServiceNames(serviceNames ...string) *common.CommandExecutor {
	return &common.CommandExecutor{Command: up.Command + helpers.ServiceName(serviceNames...)}
}
