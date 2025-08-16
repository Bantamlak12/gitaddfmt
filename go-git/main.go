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

		cmd := exec.Command("go", "fmt", "./...")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Formatting failed. Check the following errors: \n%s", output)
			os.Exit(1)
		}
	}

	fmt.Println("Go code formated.")
	args := os.Args[1:]
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}

	if os.Args[1] == "add" && len(os.Args) > 2 {
		fmt.Println("Go code staged to git.")
	}
}
