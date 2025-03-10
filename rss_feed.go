package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
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

	request.Header.Set("User-Agent", "gator")
	response, err := http.DefaultClient.Do(request)
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

	for _, itm := range feed.Channel.Item {
		itm.Description = html.UnescapeString(itm.Description)
		itm.Title = html.UnescapeString(itm.Title)
	}

	return &feed, nil
}
