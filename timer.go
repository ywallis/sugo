package main

import (
	"fmt"
	"strings"
	"time"
)

func timer(duration time.Duration, cycleType string, iteration int) {

	// Weird ass screen clear
	fmt.Print("\033[H\033[2J")
	totalSeconds := int(duration.Seconds())

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	start := time.Now()
	for range ticker.C {
		elapsed := int(time.Since(start).Seconds())
		remaining := totalSeconds - elapsed

		minutes := remaining / 60
		seconds := remaining % 60

		barLength := 30
		filled := barLength * elapsed / totalSeconds
		empty := barLength - filled
		var color string
		switch cycleType {
		case "Work":
			color = "\033[31m" // Red
		default:
			color = "\033[32m" // Green
		}

		resetColor := "\033[0m"

		bar := fmt.Sprintf(
			"%s%s%s%s",
			color,
			strings.Repeat("█", filled),
			resetColor,
			strings.Repeat("░", empty),
		)

		fmt.Printf(
			"\r%s %d - %02d:%02d [%s]",
			cycleType,
			iteration,
			minutes,
			seconds,
			bar,
		)

		if remaining < 0 {
			break
		}
	}
	fmt.Printf("\n%s done!\n", cycleType)
}
