package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func appendString(cmd string) string {
	return " " + cmd
}

type KeyValueParameters struct {
	Key   string
	Value string
}

func Option(cmd string) string {
	return appendString(cmd)
}

func Command(cmd string) string {
	return appendString(cmd)
}

func ServiceName(services ...string) string {
	var cmd string
	if len(services) != 0 {
		for _, service := range services {
			cmd += appendString(service)
		}
	}

	return cmd
}

// TODO Test it on Windows and MacOs
func GeneralExec(command string, privileges bool) {
	shell := detectShell()

	if shell == "" {
		fmt.Println("No compatible shell found!")
		return
	}

	if privileges && runtime.GOOS != "windows" {
		command = "sudo " + command
	}

	cmd := exec.Command(shell, "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error in execution of command:", err)
	}
}

func detectShell() string {
	switch runtime.GOOS {
	case "windows":
		return "powershell"
	default:
		if _, err := exec.LookPath("bash"); err == nil {
			return "bash"
		}
		if _, err := exec.LookPath("sh"); err == nil {
			return "sh"
		}
	}
	return ""
}
