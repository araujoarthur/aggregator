package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd command) error {

	f, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(f)
	return nil
}
