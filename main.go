package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type config struct {
	workCycleLength   time.Duration
	smallBreakLength  time.Duration
	longBreakLength   time.Duration
	cycles            int
	barLength         int
	confirmToContinue bool
}

func main() {
	fmt.Print("\033[?25l") // hide cursor

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c                      // Wait for signal
		fmt.Println("\033[?25h") // show cursor
		clearScreen()
		os.Exit(0) // Or use return if you prefer not to force exit
	}()

	workCycleLength := flag.Int("work", 25, "The minutes in a work cycle")
	smallBreakLength := flag.Int("break", 5, "The minutes in a small break")
	longBreakLength := flag.Int("long", 25, "The minutes in a long break")
	cycles := flag.Int("cycles", 3, "The amount of work/break cycles")
	barLength := flag.Int("bar", 30, "The length of the progress bar")
	confirmToContinue := flag.Bool("confirm", true, "Confirm to continur to next cycle")
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
		workCycleLength:   time.Minute * time.Duration(*workCycleLength),
		smallBreakLength:  time.Minute * time.Duration(*smallBreakLength),
		longBreakLength:   time.Minute * time.Duration(*longBreakLength),
		cycles:            *cycles,
		barLength:         *barLength,
		confirmToContinue: *confirmToContinue,
	}

	for i := range cfg.cycles {

		timer(cfg, "Work", i+1)
		timer(cfg, "Break", i+1)
	}
	timer(cfg, "Long break", 0)
	printCenter("Press enter to continue.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println("\033[?25h") // show cursor
	clearScreen()
}
