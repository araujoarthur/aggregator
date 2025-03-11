package main

import (
	"fmt"
	"log"
	"time"
)

func handleAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs> (1h, 1m, 1s)", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("could not parse duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
