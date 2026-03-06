package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("agg expects one argument, timeBetweenRequests. Example 1m0s\n")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		fmt.Println()
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
		fmt.Println()
		fmt.Println("===============================================================")
	}
}
