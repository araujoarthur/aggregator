package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/araujoarthur/aggregator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	usr, err := s.DbQueries.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("something went wrong retrieving user data: %w", err)
	}

	createFeedParameters := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      cmd.Args[0],
		Url:       sql.NullString{String: cmd.Args[1], Valid: true},
		UserID:    uuid.NullUUID{UUID: usr.ID, Valid: true},
	}

	_, err = s.DbQueries.CreateFeed(context.Background(), createFeedParameters)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Println("Feed created successfully!")
	return nil
}
