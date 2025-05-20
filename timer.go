package main

import (
	"fmt"
	"time"
)

func timer(duration time.Duration, cycleType string, iteration int) {

	// Weird ass screen clear
	// fmt.Print("\033[H\033[2J")
	end := time.Now().Add(duration)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		remaining := time.Until(end)
		if iteration == 0 {
			fmt.Printf("\r%s - %02d:%02d", cycleType, int(remaining.Minutes()), int(remaining.Seconds())%60)
		} else {
			fmt.Printf("\r%s %d - %02d:%02d", cycleType, iteration, int(remaining.Minutes()), int(remaining.Seconds())%60)
		}
		if remaining < 0 {
			break
		}
	}
	fmt.Printf("\n%s done!\n", cycleType)
}
