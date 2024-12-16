package main

import (
	"fmt"
	"github.com/rojack96/gocker"
)

func main() {
	g := gocker.Gocker{}
	res := g.Compose().FileName("compose.yml").Up().GetCommand()
	fmt.Println(res)
}
