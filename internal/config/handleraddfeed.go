package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Darkpowercross/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {

	if len(cmd.Args) != 2 {
		return fmt.Errorf("not enough arguments")
	}

	user, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {

		return fmt.Errorf("failed to fetch user: %w", err)
	}

	feedname := cmd.Args[0]
	feedurl := cmd.Args[1]

	fmt.Println(feedname)
	fmt.Println(feedurl)

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      sql.NullString{String: feedname, Valid: true},
		Url:       sql.NullString{String: feedurl, Valid: true},
		UserID:    user.ID,
	})

	if err != nil {
		return fmt.Errorf("error creating feed follow")
	}

	currentuserid, err := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("please login")
	}

	s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentuserid.ID,
		FeedID:    feed.ID,
	})

	return nil

}
