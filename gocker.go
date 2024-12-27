package gocker

import (
	"github.com/rojack96/gocker/commands/compose"
	"github.com/rojack96/gocker/commands/run"
	"os"
	"runtime"
	"strings"
)

const (
	docker = "docker"
)

func Run() *run.Run {
	return &run.Run{Command: docker + " run"}
}

// Compose - Define and run multi-container applications with Docker
func Compose() *compose.Compose {
	return compose.New(docker + " compose")
}

func ConcatCommands(commands ...string) string {
	if len(commands) == 0 {
		return ""
	}

	var separator string
	if runtime.GOOS == "windows" {
		shell := os.Getenv("SHELL")
		if strings.Contains(shell, "powershell") {
			separator = "; "
		} else {
			separator = " && "
		}
	} else {
		separator = " && "
	}

	return strings.Join(commands, separator)
}
