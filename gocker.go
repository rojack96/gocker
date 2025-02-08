package gocker

import (
	"github.com/rojack96/gocker/commands/compose"
	"github.com/rojack96/gocker/commands/run"
	"github.com/rojack96/gocker/commands/system"
)

const (
	docker = "docker"
)

//  Common Commands

// Run - Create and run a new container from an image.
func Run() *run.Run {
	return run.New(docker + " run")
}

//	Management Commands

// Compose - Define and run multi-container applications with Docker.
func Compose() *compose.Compose {
	return compose.New(docker + " compose")
}

func System() *system.System {
	return system.New(docker + " system")
}
