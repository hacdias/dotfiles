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
	"uno"

	"golang.org/x/oauth2"
)

type swarm struct {
	config *config
	oauth2 *oauth2.Config
	token  *oauth2.Token
}

func newSwarm(c *config) (*swarm, error) {
	t := &swarm{
		config: c,
		oauth2: &oauth2.Config{
			ClientID:     c.Client,
			ClientSecret: c.Secret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://foursquare.com/oauth2/authorize",
				TokenURL: "https://foursquare.com/oauth2/access_token",
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

func (t *swarm) interactiveLogin(port int) error {
	token, err := uno.InteractiveLogin(t.oauth2, port)
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

func (t *swarm) fetch(ctx context.Context, page int, start time.Time, end time.Time) (swarmCheckins, bool, error) {
	limit := 250
	u, err := url.Parse("https://api.foursquare.com/v2/users/self/checkins")
	if err != nil {
		return nil, false, err
	}

	q := u.Query()
	q.Set("sort", "newestfirst")
	q.Set("limit", strconv.Itoa(limit))
	q.Set("offset", strconv.Itoa(limit*page))
	q.Set("oauth_token", t.token.AccessToken)
	q.Set("v", "20231010")

	if !start.IsZero() {
		q.Set("afterTimestamp", strconv.FormatInt(start.Unix(), 10))
	}

	if !end.IsZero() {
		q.Set("beforeTimestamp", strconv.FormatInt(end.Unix(), 10))
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
	req.Header.Set("Accept", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer res.Body.Close()

	var swarmResponse *swarmOutput

	err = json.NewDecoder(res.Body).Decode(&swarmResponse)
	if err != nil {
		return nil, false, err
	}

	if swarmResponse.Meta.Code != 200 {
		return nil, false, fmt.Errorf("response status is not 200, is %d", swarmResponse.Meta.Code)
	}

	return swarmResponse.Response.Checkins.Items, len(swarmResponse.Response.Checkins.Items) != 0, nil
}

func (t *swarm) fetchBetween(ctx context.Context, start, end time.Time) error {
	if t.token == nil {
		return errors.New("token is not set")
	}

	var (
		err      error
		page     = 0
		checkins swarmCheckins
		hasMore  = true
	)

	for hasMore {
		var localCheckins swarmCheckins
		localCheckins, hasMore, err = t.fetch(ctx, page, start, end)
		if err != nil {
			return err
		}

		checkins = append(checkins, localCheckins...)
		time.Sleep(time.Millisecond * 500) // Beware of those rate limitings.
		page++
	}

	err = t.saveHistory(checkins)
	if err != nil {
		return err
	}

	return t.updateToken()
}

func (t *swarm) saveHistory(history swarmCheckins) error {
	history.sort()

	if len(history) == 0 {
		return nil
	}

	h := map[int]map[time.Month]swarmCheckins{}

	for i := range history {
		year := history[i].CreatedAt().Year()
		month := history[i].CreatedAt().Month()

		if _, ok := h[year]; !ok {
			h[year] = map[time.Month]swarmCheckins{}
		}

		if _, ok := h[year][month]; !ok {
			h[year][month] = swarmCheckins{}
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

func (t *swarm) updateToken() error {
	raw, err := json.MarshalIndent(t.token, "", "  ")
	if err != nil {
		return nil
	}

	return os.WriteFile(t.config.Token, raw, 0644)
}
