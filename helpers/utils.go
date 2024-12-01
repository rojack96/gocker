package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func AppendString(cmd string) string {
	return " " + cmd
}

// TODO Test it on Windows and MacOs
func GeneralExec(command string) {
	shell := detectShell()

	if shell == "" {
		fmt.Println("No compatible shell found!")
		return
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
		return "cmd"
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
