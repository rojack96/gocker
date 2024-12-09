package compose

import (
	"fmt"
	"gocker/helpers"
)

type Up struct {
	Command string
}

func (up *Up) AbortOnContainerExit() *Up {
	up.Command = " --abort-on-container-exit"
	return up
}

func (up *Up) AbortOnContainerFailure() *Up {
	up.Command = " --abort-on-container-failure"
	return up
}

func (up *Up) AlwaysRecreateDeps() *Up {
	up.Command = " --always-recreate-deps"
	return up
}

func (up *Up) Attach(services []string) *Up {
	for _, service := range services {
		up.Command += " --attach " + service
	}
	return up
}

func (up *Up) AttachDependencies() *Up {
	up.Command = " --attach-dependencies"
	return up
}

func (up *Up) Build() *Up {
	up.Command = " --build"
	return up
}

func (up *Up) Detach() *Up {
	up.Command = " -d"
	return up
}

func (up *Up) DryRun() *Up {
	up.Command = " --dry-run"
	return up
}

func (up *Up) ExitCodeFrom(service string) *Up {
	up.Command += " --exit-code-from " + service
	return up
}

func (up *Up) ForceRecreate() *Up {
	up.Command = " --force-recreate"
	return up
}

func (up *Up) NoAttach(services []string) *Up {
	for _, service := range services {
		up.Command += " --no-attach " + service
	}
	return up
}

func (up *Up) NoBuild() *Up {
	up.Command = " --no-build"
	return up
}

func (up *Up) NoColor() *Up {
	up.Command = " --no-color"
	return up
}

func (up *Up) NoDeps() *Up {
	up.Command = " --no-deps"
	return up
}

func (up *Up) NoLogPrefix() *Up {
	up.Command = " --no-log-prefix"
	return up
}

func (up *Up) NoRecreate() *Up {
	up.Command = " --no-recreate"
	return up
}

func (up *Up) NoStart() *Up {
	up.Command = " --no-start"
	return up
}

func (up *Up) Pull(pullPolicy string) *Up {
	up.Command += " --pull " + pullPolicy
	return up
}

func (up *Up) QuietPull() *Up {
	up.Command = " --quiet-pull"
	return up
}

func (up *Up) RemoveOrphans() *Up {
	up.Command = " --remove-orphans"
	return up
}

func (up *Up) RenewAnonVolumes() *Up {
	up.Command = " -V"
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
	up.Command = " --timestamps"
	return up
}

func (up *Up) Wait() *Up {
	up.Command = " --wait"
	return up
}

func (up *Up) WaitTimeout(seconds int) *Up {
	up.Command += " --wait-timeout " + fmt.Sprintf("%d", seconds)
	return up

}

func (up *Up) Watch() *Up {
	up.Command = " -w"
	return up
}

func (up *Up) ServiceName(serviceName ...string) *Up {
	up.Command = helpers.ServiceName(serviceName...)
	return up
}

func (up *Up) GetCommand() string {
	return up.Command
}

func (up *Up) Exec() {
	helpers.GeneralExec(up.Command)
}
