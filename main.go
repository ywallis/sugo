package main

import (
	"fmt"
	"time"
)

const secondsInMinutes int = 3

type config struct {
	work_cycle_length  time.Duration
	small_break_length time.Duration
	long_break_length  time.Duration
	cycles             int
}

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

func main() {
	cfg := config{
		work_cycle_length:  time.Second * 3,
		small_break_length: time.Second,
		long_break_length:  time.Second * 3,
		cycles:             3,
	}
	for i := range cfg.cycles {

		timer(cfg.work_cycle_length, "Work", i+1)
		timer(cfg.small_break_length, "Small break", i+1)
	}
	fmt.Println("Long break")
	timer(cfg.long_break_length, "Long break", 0)
}
