package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/araujoarthur/aggregator/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command, user database.User) error {

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

	if follows {
		return fmt.Errorf("user already follows the feed %s [%s]", feed.Name, feed.Url.String)
	}

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	result, err := s.DbQueries.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("error when creating a new follow record: %w", err)
	}

	fmt.Printf("Feed %s followed by %s successfully!\n", result.FeedName, result.UserName)

	return nil
}
