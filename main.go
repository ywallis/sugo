package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

type config struct {
	work_cycle_length  time.Duration
	small_break_length time.Duration
	long_break_length  time.Duration
	cycles             int
}

func main() {
	cfg := config{}
	cfg.work_cycle_length = time.Minute * 1
	cfg.small_break_length = time.Minute * 5
	cfg.long_break_length = time.Minute * 25

	args := os.Args[1:]
	
	if len(args) > 1 {
		log.Fatalln("A single cycle number argument is accepted")
	}
	// Default config
	
	if len(args) == 0 {
		cfg.cycles = 3
	} else {
		cycles, err := strconv.Atoi(args[0]) 
		if err != nil {
			log.Fatal("Enter an interger as a cycle amount")
		}
		cfg.cycles = cycles 
	}

	for i := range cfg.cycles {

		timer(cfg.work_cycle_length, "Work", i+1)
		timer(cfg.small_break_length, "Break", i+1)
	}
	timer(cfg.long_break_length, "Long break", 0)
}
