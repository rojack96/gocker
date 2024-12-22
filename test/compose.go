package main

import (
	"fmt"
	"github.com/rojack96/gocker"
)

func main() {
	g := gocker.Compose()
	// TODO write tests
	res := g.FileName().Attach().DryRun().Index(5).ServiceName("hello-world").GetCommand()
	fmt.Println(res)
}
