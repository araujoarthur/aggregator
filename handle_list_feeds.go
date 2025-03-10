package main

import (
	"context"
	"fmt"
)

func handleListFeeds(s *state, cmd command) error {
	feeds, err := s.DbQueries.GetFeedsWithUserInfo(context.Background())
	if err != nil {
		return fmt.Errorf("something went wrong returning feeds: %w", err)
	}

	for _, feed := range feeds {
		if feed.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", feed.Name)
			continue
		}
		fmt.Printf("* %s \n", feed.Name)
	}

	return nil
}
