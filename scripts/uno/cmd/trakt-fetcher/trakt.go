package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
	"uno"

	"golang.org/x/oauth2"
)

type trakt struct {
	config *config
	oauth2 *oauth2.Config
	token  *oauth2.Token
}

func newTrakt(c *config) (*trakt, error) {
	t := &trakt{
		config: c,
		oauth2: &oauth2.Config{
			// No client secret: login uses PKCE.
			ClientID: c.Client,
			Endpoint: oauth2.Endpoint{
				AuthURL:   "https://auth.trakt.tv/oauth/authorize",
				TokenURL:  "https://auth.trakt.tv/oauth/token",
				AuthStyle: oauth2.AuthStyleInParams,
			},
		},
	}

	raw, err := os.ReadFile(c.Token)
	if err == nil {
		err = json.Unmarshal(raw, &t.token)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func (t *trakt) interactiveLogin(port int) error {
	token, err := uno.InteractiveLogin(t.oauth2, port)
	if err != nil {
		return err
	}

	t.token = token
	err = t.updateToken()
	if err != nil {
		return err
	}

	fmt.Println("Token updated.")
	return nil
}

func (t *trakt) fetch(ctx context.Context, client *http.Client, page int, start time.Time, end time.Time) (traktHistory, bool, error) {
	limit := 100
	u, err := url.Parse("https://api.trakt.tv/sync/history")
	if err != nil {
		return nil, false, err
	}

	q := u.Query()
	q.Set("extended", "full")
	q.Set("limit", strconv.Itoa(limit))
	q.Set("page", strconv.Itoa(page))

	if !start.IsZero() {
		q.Set("start_at", start.Format(time.RFC3339Nano))
	}

	if !end.IsZero() {
		q.Set("end_at", end.Format(time.RFC3339Nano))
	}

	u.RawQuery = q.Encode()

	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, false, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "trakt-fetcher")
	req.Header.Set("trakt-api-key", t.config.Client)
	req.Header.Set("trakt-api-version", "2")

	res, err := client.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("response status is %s", res.Status)
	}

	currentPage, err := strconv.Atoi(res.Header.Get("X-Pagination-Page"))
	if err != nil {
		return nil, false, err
	}

	totalPages, err := strconv.Atoi(res.Header.Get("X-Pagination-Page-Count"))
	if err != nil {
		return nil, false, err
	}

	var history traktHistory

	err = json.NewDecoder(res.Body).Decode(&history)
	if err != nil {
		return nil, false, err
	}

	return history, currentPage < totalPages, nil
}

func (t *trakt) fetchBetween(ctx context.Context, start, end time.Time) error {
	if t.token == nil {
		return errors.New("token is not set")
	}

	var (
		err         error
		page        = 1
		history     traktHistory
		historyPage traktHistory
		hasMore     = true
	)

	// Refresh tokens are single-use: save a refreshed token before anything else
	// can fail, and use a static token so nothing refreshes without being saved.
	t.token, err = t.oauth2.TokenSource(ctx, t.token).Token()
	if err != nil {
		return err
	}

	err = t.updateToken()
	if err != nil {
		return err
	}

	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(t.token))

	for hasMore {
		historyPage, hasMore, err = t.fetch(ctx, client, page, start, end)
		if err != nil {
			return err
		}

		history = append(history, historyPage...)
		time.Sleep(time.Millisecond * 500) // Beware of those rate limitings.
		page++
	}

	history.sort()
	return uno.SaveByMonth(t.config.Output, history, func(i traktHistoryItem) time.Time { return i.WatchedAt })
}

func (t *trakt) updateToken() error {
	raw, err := json.MarshalIndent(t.token, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(t.config.Token, raw, 0600)
}
