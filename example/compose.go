package main

import (
	"fmt"
	"github.com/rojack96/gocker"
)

func main() {
	compose := gocker.Compose()

	result := compose.FileName("compose.yml").Up().Build().ServiceNames()

	fmt.Println(result.GetCommand())

	result.Exec(false)
}
