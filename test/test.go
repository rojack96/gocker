package main

import (
	"gocker"
)

func main() {
	g := gocker.Gocker{}
	g.Compose().FileName("compose.yml").Up().Exec()
}
