package config

import (
	"context"
	"fmt"
)

func Handlerfeeds(s *State, cmd Command) error {
	feeds, err := s.DB.ListFeed(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds")
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name: %s, Feed URL, %s, Created by: %s \n", feed.FeedName.String, feed.FeedUrl.String, feed.UserName)
	}

	return nil
}
