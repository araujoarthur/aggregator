package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/araujoarthur/aggregator/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	request.Header.Set("User-Agent", "gator")
	response, err := client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("something went wrong requesting data: %w", err)
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response body: %w", err)
	}

	var feed RSSFeed
	if err = xml.Unmarshal(data, &feed); err != nil {
		return nil, fmt.Errorf("error unmarshalling the response body: %w", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i, itm := range feed.Channel.Item {
		itm.Description = html.UnescapeString(itm.Description)
		itm.Title = html.UnescapeString(itm.Title)
		feed.Channel.Item[i] = itm
	}

	return &feed, nil
}

func scrapeFeeds(s *state) {
	feed, err := s.DbQueries.GetNextFeedToFetch(context.Background())

	if err != nil {
		log.Println("could not retrieve next feed", err)
		return
	}

	log.Println("Found a feed to fetch: ", feed.Name)
	scrapeFeed(s.DbQueries, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	err := db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true}, ID: feed.ID})
	if err != nil {
		log.Println("could not mark feed as fetched", err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url.String)
	if err != nil {
		log.Printf("could not fetch feed %s's data: %v", feed.Name, err)
		return
	}

	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}

			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
