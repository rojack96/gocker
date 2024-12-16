package gocker

import (
	"github.com/rojack96/gocker/commands/compose"
	"github.com/rojack96/gocker/commands/run"
)

const (
	Docker = "docker"
)

type Gocker struct{}

func (g *Gocker) Run() *run.Run {
	return &run.Run{Command: Docker + " run"}
}

func (g *Gocker) Compose() *compose.Compose {
	return &compose.Compose{Command: Docker + " compose"}
}
