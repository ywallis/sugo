package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"
)

func clearScreen() {
	cmd := exec.Command("clear") // Use "cls" on Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func verticalAlign() {

	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}
	verticalPadding := (height / 2) - 1

	clearScreen()

	// Print blank lines for vertical centering
	for range verticalPadding {
		fmt.Println()
	}
}
func printCenter(text string) {

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	lines := strings.SplitSeq(text, "\n")

	for line := range lines {
		lineWidth := utf8.RuneCountInString(line)
		horizPad := max((width-lineWidth)/2, 0)
		fmt.Printf("\r%s%s", strings.Repeat(" ", horizPad), line)
	}
}
