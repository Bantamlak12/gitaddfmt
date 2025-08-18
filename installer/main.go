package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func startOllamaBackground() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "Start-Service", "Ollama")
	case "linux":
		cmd = exec.Command("systemctl", "start", "ollama")
	case "darwin":
		cmd = exec.Command("brew", "services", "start", "ollama")
	default:
		fmt.Println("Unsupported OS for starting Ollama service.")
		os.Exit(1)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start Ollama:", err)
		os.Exit(1)
	}

	fmt.Println("Ollama service started successfully in the background (PID:", cmd.Process.Pid, ")")
}

func pullOllamaModel() {
	cmd := exec.Command("ollama", "pull", "gpt-oss:20b")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to pull gpt-oss:20b model:", err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Installing go-git...")

	cmd := exec.Command("go", "install", "github.com/Bantamlak12/gitaddfmt/go-git@latest")
	// cmd := exec.Command("go", "install", "./go-git")
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

	// Check if ollama is installed
	_, err := exec.LookPath("ollama")
	if err != nil {
		fmt.Println("Downloading and installing Ollama... this may take a few minutes.")

		var installCmd *exec.Cmd
		switch runtime.GOOS {
		case "windows":
			installCmd = exec.Command("powershell", "Invoke-WebRequest https://ollama.com/download/OllamaSetup.exe -OutFile OllamaSetup.exe; Start-Process -FilePath OllamaSetup.exe -Wait; Remove-Item OllamaSetup.exe")
		case "linux", "darwin":
			installCmd = exec.Command("bash", "-c", "set -e -o pipefail && curl -fsSL https://ollama.com/install.sh | sh")
		default:
			fmt.Println("Unsupported OS for ollama installation.")
			os.Exit(1)
		}

		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		if err := installCmd.Run(); err != nil {
			fmt.Println("Failed to install ollama:", err)
			os.Exit(1)
		}
		fmt.Println("ollama installed successfully.")
	}

	startOllamaBackground()
	pullOllamaModel()
}
