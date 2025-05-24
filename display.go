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

func centerText(text string) {

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	lines := strings.Split(text, "\n")
	verticalPadding := (height - len(lines)) / 2

	clearScreen()

	fmt.Print("\033[?25l") // hide cursor
	// defer fmt.Print("\033[?25h")
	// Print blank lines for vertical centering
	for range verticalPadding {
		fmt.Println()
	}

	for _, line := range lines {
		lineWidth := utf8.RuneCountInString(line)
		horizPad := max((width - lineWidth) / 2, 0)
		fmt.Printf("%s%s\n", strings.Repeat(" ", horizPad), line)
	}
}
