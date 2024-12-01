package gocker

const (
	DockerCompose = "docker compose"
	DockerRun     = "docker run"
)

type Gocker struct{}

func (g *Gocker) Run() *Run {
	return &Run{Command: DockerRun}
}

func (g *Gocker) Compose() *Compose {
	return &Compose{Command: DockerCompose}
}
