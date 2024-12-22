package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

type KeyValueParameters struct {
	Key   string
	Value string
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

func appendString(cmd string) string {
	return " " + cmd
}

func Option(cmd string) string {
	return appendString(cmd)
}

func String(cmd, value string) string {
	return appendString(cmd) + appendString(value)
}

func Int(cmd string, value int) string {
	return appendString(cmd) + appendString(strconv.Itoa(value))
}

func StringArray(cmd string, array ...string) string {
	var result string
	if len(array) != 0 {
		for _, arr := range array {
			result += appendString(cmd) + appendString(arr)
		}
	}

	return result
}

func Command(cmd string) string {
	return appendString(cmd)
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
