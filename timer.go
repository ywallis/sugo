package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
)

func timer(config config, cycleType string, iteration int) {

	var duration time.Duration
	switch cycleType {
	case "Work":
		duration = config.workCycleLength
	case "Break":
		duration = config.smallBreakLength
	case "Long break":
		duration = config.longBreakLength
	}

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

		barLength := config.barLength
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
		if cycleType == "Long break" {
			centerText(fmt.Sprintf(
				"%s - %02d:%02d [%s]",
				cycleType,
				minutes,
				seconds,
				bar))
		} else {
			centerText(fmt.Sprintf(
				"%s %d of %d - %02d:%02d [%s]",
				cycleType,
				iteration,
				config.cycles,
				minutes,
				seconds,
				bar))
		}

		if remaining < 0 {
			break
		}
	}
	centerText(fmt.Sprintf("\n%s done!\n", cycleType))
	if err := beeep.Alert(cycleType, "Done!", "assets/sugo.png"); err != nil {
		log.Fatalln("Could not send notification:", err)
	}
}
