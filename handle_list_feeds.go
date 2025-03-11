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
		fmt.Printf("* %s [%s]: %s\n", feed.Feed.Name, feed.User.Name, feed.Feed.Url.String)
	}

	return nil
}
