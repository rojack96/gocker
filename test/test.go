package main

import (
	"fmt"
	"gocker"
)

func main() {
	g := gocker.Gocker{}
	test := g.Compose().FileName("docker-compose-dev.yml").Up().Build().Detach()
	fmt.Println(test.Command)
}
