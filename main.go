package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "gui" {
		runGUI()
	} else {
		runCLI()
	}
}

func runCLI() {
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

func runGUI() {
	a := app.New()
	w := a.NewWindow("MP4 to Frames")

	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Input file (.mp4)")

	outputEntry := widget.NewEntry()
	outputEntry.SetPlaceHolder("Output directory (default: frames)")

	selectFileBtn := widget.NewButton("Select File", func() {
		dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if r != nil {
				inputEntry.SetText(r.URI().Path())
			}
		}, w)
	})

	extractBtn := widget.NewButton("Extract Frames", func() {
		inputFile := inputEntry.Text
		outputDir := outputEntry.Text
		if outputDir == "" {
			outputDir = "frames"
		}

		if _, err := os.Stat(inputFile); os.IsNotExist(err) {
			dialog.ShowError(fmt.Errorf("input file %s does not exist", inputFile), w)
			return
		}

		// Ensure output directory exists
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			err := os.Mkdir(outputDir, 0755)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to create output directory: %s", err), w)
				return
			}
		}

		outputPattern := filepath.Join(outputDir, "frame_%04d.png")

		cmd := exec.Command("ffmpeg", "-i", inputFile, outputPattern)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to execute ffmpeg command: %s", err), w)
			return
		}

		dialog.ShowInformation("Success", "Frames extracted successfully to "+outputDir, w)
	})

	w.SetContent(container.NewVBox(
		inputEntry,
		selectFileBtn,
		outputEntry,
		extractBtn,
	))

	w.ShowAndRun()
}
