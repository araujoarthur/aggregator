package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/araujoarthur/aggregator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	createFeedParameters := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      cmd.Args[0],
		Url:       sql.NullString{String: cmd.Args[1], Valid: true},
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
	}

	_, err := s.DbQueries.CreateFeed(context.Background(), createFeedParameters)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Println("Feed created successfully!")

	err = s.RegisteredCommands.run(s, command{Name: "follow", Args: []string{cmd.Args[1]}})
	if err != nil {
		return fmt.Errorf("could not automatically follow the feed: %w", err)
	}

	return nil
}
