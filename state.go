package main

import (
	"context"

	"github.com/araujoarthur/aggregator/internal/config"
	"github.com/araujoarthur/aggregator/internal/database"
)

type state struct {
	Config             *config.Config
	DbQueries          *database.Queries
	RegisteredCommands *commands
}

func (s *state) UserExists(name string) bool {
	if _, err := s.DbQueries.GetUser(context.Background(), name); err == nil { // If no error happens, the user exists.
		return true
	}
	return false
}
