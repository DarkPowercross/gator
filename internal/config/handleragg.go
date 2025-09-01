package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Darkpowercross/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage gator agg <time>")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapefeeds(s)
	}
}

func scrapefeeds(s *State) error {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("no feeds to fetch")
	}

	s.DB.MarkFeedFetched(context.Background(), feed.ID)

	feeditem, err := fetchFeed(context.Background(), feed.Url.String)
	if err != nil {
		return fmt.Errorf("error with url")
	}

	for _, item := range feeditem.Channel.Item {
		fmt.Printf("Collecting feed for: %v\n", item.Title)
		_, err = s.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{
				String: item.Title,
				Valid:  item.Title != "",
			},
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			FeedID: feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}

	return nil

}
