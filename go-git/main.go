package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go git add <files>")
		os.Exit(1)
	}

	if os.Args[1] == "add" {
		fmt.Println("Formatting Go code...")
		if err := exec.Command("go", "fmt", "./...").Run(); err != nil {
			fmt.Println("Formatting failed:", err)
			os.Exit(1)
		}
	}

	args := os.Args[1:]
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
