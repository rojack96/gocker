package main

import (
	"github.com/rojack96/gocker"
)

func main() {
	g := gocker.Gocker{}
	g.Compose().FileName("compose.yml").Up().Exec()

	g.Compose().Up().AbortOnContainerExit()
}
