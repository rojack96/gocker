package gocker

import (
	"github.com/rojack96/gocker/commands/compose"
	"github.com/rojack96/gocker/commands/run"
)

const (
	docker = "docker"
)

// Run - Create and run a new container from an image.
func Run() *run.Run {
	return run.New(docker + " run")
}

// Compose - Define and run multi-container applications with Docker.
func Compose() *compose.Compose {
	return compose.New(docker + " compose")
}
