package gocker

import (
	"github.com/rojack96/gocker/commands/compose"
	"github.com/rojack96/gocker/commands/run"
)

const (
	docker = "docker"
)

type Gocker struct{}

func (g *Gocker) Run() *run.Run {
	return &run.Run{Command: docker + " run"}
}

// Compose - Define and run multi-container applications with Docker
func (g *Gocker) Compose() *compose.Compose {
	return &compose.Compose{Command: docker + " compose"}
}
