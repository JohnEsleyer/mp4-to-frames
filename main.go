package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mp4-to-frames <input-file> [output-directory]")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputDir := "frames"
	if len(os.Args) == 3 {
		outputDir = os.Args[2]
	}

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("Input file %s does not exist\n", inputFile)
		os.Exit(1)
	}

	// Ensure output directory exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, 0755)
		if err != nil {
			fmt.Printf("Failed to create output directory: %s\n", err)
			os.Exit(1)
		}
	}

	outputPattern := filepath.Join(outputDir, "frame_%04d.png")

	cmd := exec.Command("ffmpeg", "-i", inputFile, outputPattern)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running ffmpeg command:", strings.Join(cmd.Args, " "))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to execute ffmpeg command: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Frames extracted successfully to", outputDir)
}
