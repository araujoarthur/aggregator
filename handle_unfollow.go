package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/araujoarthur/aggregator/internal/database"
)

func handleUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.DbQueries.GetFeedByURL(context.Background(), sql.NullString{String: cmd.Args[0], Valid: true})
	if err != nil {
		return fmt.Errorf("failed to fetch feed data: %w", err)
	}

	follows, err := s.DbQueries.UserFollowsFeed(context.Background(), database.UserFollowsFeedParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return fmt.Errorf("failed to fetch follow data: %w", err)
	}

	if !follows {
		return fmt.Errorf("user %s does not follow the feed %s [%s]", user.Name, feed.Name, feed.Url.String)
	}

	unfollowParams := database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.DbQueries.UnfollowFeed(context.Background(), unfollowParams)
	if err != nil {
		return fmt.Errorf("error when creating a new follow record: %w", err)
	}

	fmt.Printf("User %s successfuly unfollowed the feed %q!\n", user.Name, feed.Name)

	return nil
}
