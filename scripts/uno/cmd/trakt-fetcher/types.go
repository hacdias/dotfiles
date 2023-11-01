package main

import (
	"sort"
	"time"
)

// https://trakt.docs.apiary.io/#introduction/standard-media-objects

type traktGeneric struct {
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	IDs     traktIDs `json:"ids"`
	Runtime int      `json:"runtime,omitempty"`
	Country string   `json:"country,omitempty"`
	Genres  []string `json:"genres,omitempty"`
}

type traktMovie traktGeneric

type traktShow traktGeneric

type traktEpisode struct {
	Title  string   `json:"title"`
	Season int      `json:"season"`
	Number int      `json:"number"`
	IDs    traktIDs `json:"ids"`
}

type traktIDs struct {
	Trakt int    `json:"trakt"`
	IMDb  string `json:"imdb"`
	TMDb  int    `json:"tmdb"`
	Slug  string `json:"slug,omitempty"`
	TVDb  int    `json:"tvdb,omitempty"`
}

type traktHistoryItem struct {
	ID        int64         `json:"id"`
	WatchedAt time.Time     `json:"watched_at"`
	Action    string        `json:"action"`
	Type      string        `json:"type"`
	Movie     *traktMovie   `json:"movie,omitempty"`
	Episode   *traktEpisode `json:"episode,omitempty"`
	Show      *traktShow    `json:"show,omitempty"`
}

type traktHistory []traktHistoryItem

func (h traktHistory) sort() {
	sort.SliceStable(h, func(i, j int) bool {
		return h[i].WatchedAt.After(h[j].WatchedAt)
	})
}
