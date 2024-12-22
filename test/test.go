package main

import (
	"fmt"
	"github.com/rojack96/gocker"
)

func main() {
	g := gocker.Compose()
	//res := g.Compose().FileName("compose.yml").Up().GetCommand()

	res := g.FileName().Attach().DryRun().Index(5).ServiceName().GetCommand()
	fmt.Println(res)
}
