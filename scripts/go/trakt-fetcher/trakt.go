package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
			ClientID:     c.Client,
			ClientSecret: c.Secret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://trakt.tv/oauth/authorize",
				TokenURL: "https://trakt.tv/oauth/token",
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
	state, err := randomString(10)
	if err != nil {
		return nil
	}

	t.oauth2.RedirectURL = fmt.Sprintf("http://localhost:%d/callback", port)

	url := t.oauth2.AuthCodeURL(state)
	fmt.Printf("Please open the following URL, authenticate, and close the tab:\n%s\n", url)

	request := make(chan *http.Request, 1)

	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request <- r
	})}

	go func() {
		_ = server.ListenAndServe()
	}()
	defer func() {
		_ = server.Shutdown(context.Background())
	}()

	r := <-request

	if s := r.URL.Query().Get("state"); s != state {
		return errors.New("state does not match")
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		return errors.New("code was empty")
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	token, err := t.oauth2.Exchange(ctx, code)
	if err != nil {
		return nil
	}

	t.token = token
	err = t.updateToken()
	if err != nil {
		return nil
	}

	fmt.Println("Token updated.")
	return nil
}

func (t *trakt) fetch(ctx context.Context, page int, start time.Time, end time.Time) (traktHistory, bool, error) {
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

	t.oauth2.Client(ctx, t.token)

	httpClient := t.oauth2.Client(ctx, t.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("trakt-api-key", t.config.Client)
	req.Header.Set("trakt-api-version", "2")

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()

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
	var (
		err         error
		page        = 1
		history     traktHistory
		historyPage traktHistory
		hasMore     = true
	)

	for hasMore {
		historyPage, hasMore, err = t.fetch(ctx, page, start, end)
		if err != nil {
			return err
		}

		history = append(history, historyPage...)
		time.Sleep(time.Millisecond * 500) // Beware of those rate limitings.
		page++
	}

	err = t.saveHistory(history)
	if err != nil {
		return err
	}

	return t.updateToken()
}

func (t *trakt) saveHistory(history traktHistory) error {
	history.sort()

	if len(history) == 0 {
		return nil
	}

	h := map[int]map[time.Month]traktHistory{}

	for i := range history {
		year := history[i].WatchedAt.Year()
		month := history[i].WatchedAt.Month()

		if _, ok := h[year]; !ok {
			h[year] = map[time.Month]traktHistory{}
		}

		if _, ok := h[year][month]; !ok {
			h[year][month] = traktHistory{}
		}

		h[year][month] = append(h[year][month], history[i])
	}

	for year := range h {
		for month := range h[year] {
			bytes, err := json.MarshalIndent(h[year][month], "", "\t")
			if err != nil {
				return err
			}

			err = os.WriteFile(filepath.Join(t.config.Output, fmt.Sprintf("%04d-%02d.json", year, month)), bytes, 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *trakt) updateToken() error {
	raw, err := json.MarshalIndent(t.token, "", "  ")
	if err != nil {
		return nil
	}

	return os.WriteFile(t.config.Token, raw, 0644)
}
