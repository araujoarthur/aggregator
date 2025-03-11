package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/araujoarthur/aggregator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.Args) >= 1 {
		l, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			limit = 2
		} else {
			limit = int32(l)
		}
	}

	posts, err := s.DbQueries.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})

	if err != nil {
		return fmt.Errorf("error querying posts: %w", err)
	}

	for i, post := range posts {
		fmt.Printf("%d. %s\n", i, post.Title)
	}

	fmt.Println("Finished Browsing.")
	return nil
}
