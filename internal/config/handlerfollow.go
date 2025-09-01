package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Darkpowercross/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	feedurl := cmd.Args[0]
	urlfeed, err := s.DB.GetURLFeed(context.Background(), sql.NullString{String: feedurl, Valid: true})
	if err != nil {
		return fmt.Errorf("no feed created, please create new feed")
	}

	s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    urlfeed.ID,
	})

	fmt.Printf("Feed name: %v, Username: %s\n", urlfeed.Name, user.Name)

	return nil
}

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("no follows")
	}

	for _, follow := range follows {
		fmt.Printf("%s\n", follow.FeedName.String)
	}

	return nil
}

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	feedurl := cmd.Args[0]
	urlfeed, err := s.DB.GetURLFeed(context.Background(), sql.NullString{String: feedurl, Valid: true})
	if err != nil {
		return fmt.Errorf("no feed created, please create new feed")
	}
	s.DB.UnFollow(context.Background(), database.UnFollowParams{UserID: user.ID, FeedID: urlfeed.ID})
	return nil
}