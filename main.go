package main

import (
	"flag"
	"log"
	"time"
)

type config struct {
	workCycleLength  time.Duration
	smallBreakLength time.Duration
	longBreakLength  time.Duration
	cycles           int
	barLength        int
}

func main() {
	workCycleLength := flag.Int("work", 25, "The minutes in a work cycle")
	smallBreakLength := flag.Int("break", 5, "The minutes in a small break")
	longBreakLength := flag.Int("long", 25, "The minutes in a long break")
	cycles := flag.Int("cycles", 3, "The amount of work/break cycles")
	barLength := flag.Int("bar", 30, "The length of the progress bar")
	flag.Parse()

	if *workCycleLength < 1 {
		log.Fatalln("All cycles must be at least 1 min long")
	}
	if *smallBreakLength < 1 {
		log.Fatalln("All cycles must be at least 1 min long")
	}
	if *longBreakLength < 1 {
		log.Fatalln("All cycles must be at least 1 min long")
	}

	cfg := config{
		workCycleLength:  time.Minute * time.Duration(*workCycleLength),
		smallBreakLength: time.Minute * time.Duration(*smallBreakLength),
		longBreakLength:  time.Minute * time.Duration(*longBreakLength),
		cycles:           *cycles,
		barLength:        *barLength,
	}

	for i := range cfg.cycles {

		timer(cfg, "Work", i+1)
		timer(cfg, "Break", i+1)
	}
	timer(cfg, "Long break", 0)
}
