package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("Installing go-git...")

	cmd := exec.Command("go", "install", "github.com/Bantamlak12/gitaddfmt/go-git@latest")
	//cmd := exec.Command("go", "install", "./go-git")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to install go-git:", err)
		os.Exit(1)
	}
	home, _ := os.UserHomeDir()
	shell := os.Getenv("SHELL")
	var rcFilePath string

	switch runtime.GOOS {
	case "windows":
		profile := filepath.Join(home, "Documents", "PowerShell", "Microsoft.PowerShell_profile.ps1")
		f, _ := os.OpenFile(profile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		f.WriteString(`
function go {
	param($sub, $rest)
	if ($sub -eq 'git') {go-git @rest}
	else { & go.exe $sub @rest }
}
`)
		exec.Command("powershell", "-NoProfile", "-Command", "& { . $PROFILE }").Run()
		fmt.Println("Close and reopen powershell to start using 'go git add'")

	default:
		switch {
		case filepath.Base(shell) == "bash":
			rcFilePath = filepath.Join(home, ".bashrc")
		case filepath.Base(shell) == "zsh":
			rcFilePath = filepath.Join(home, ".zshrc")
		default:
			rcFilePath = filepath.Join(home, ".bashrc")
		}

		f, _ := os.OpenFile(rcFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()

		f.WriteString(`
# Go function to use 'go git' syntax instead of go-git
go() {
    if [ "$1" = "git" ]; then
        shift
        go-git "$@"
    else
        command go "$@"
    fi
}
	`)
		exec.Command("bash", "-c", "source "+rcFilePath).Run()
	}
}
