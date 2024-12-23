package src

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func Dir() {
	for {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Could not get HOME directory:", err)
			return
		}

		configPath := filepath.Join(home, ".config")

		dirEntries, err := os.ReadDir(configPath)
		if err != nil {
			fmt.Println("Error opening ~/.config:", err)
			return
		}

		directories := []string{}
		for _, entry := range dirEntries {
			if entry.IsDir() {
				directories = append(directories, entry.Name())
			}
		}

		if len(directories) == 0 {
			fmt.Println("No directories found in ~/.config.")
			return
		}

		fmt.Println("Directories in ~/.config:")
		for i, dir := range directories {
			fmt.Printf("%d: %s\n", i+1, dir)
		}

		fmt.Print("\nSelect a directory (number) or 'x' to exit: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "x" {
			fmt.Println("Exiting...")
			return
		}

		selection, err := strconv.Atoi(input)
		if err != nil || selection < 1 || selection > len(directories) {
			fmt.Println("Invalid selection.")
			return
		}

		selectedDir := filepath.Join(configPath, directories[selection-1])
		cmd := exec.Command("nvim", selectedDir)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println("Error opening nvim. ", err)
		}

	}

}
