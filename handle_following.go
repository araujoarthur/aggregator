package main

import (
	"context"
	"fmt"

	"github.com/araujoarthur/aggregator/internal/database"
)

func handleFollowing(s *state, cmd command, user database.User) error {

	follows, err := s.DbQueries.GetFollowsByUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("something went wrong returning follows for user %s: %w", user.Name, err)
	}

	for _, follow := range follows {
		fmt.Printf("* %s \n", follow.FeedName)
	}

	return nil
}
