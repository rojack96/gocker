package gocker

import (
	"github.com/rojack96/gocker/commands/compose"
	"github.com/rojack96/gocker/commands/run"
)

const (
	docker = "docker"
)

func Run() *run.Run {
	return &run.Run{Command: docker + " run"}
}

// Compose - Define and run multi-container applications with Docker
func Compose() *compose.Compose {
	return &compose.Compose{Command: docker + " compose"}
}
