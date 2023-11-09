package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type lastFm struct {
	*config
}

func newLastFm(c *config) *lastFm {
	return &lastFm{
		config: c,
	}
}

func (l *lastFm) fetchBetween(ctx context.Context, from, to time.Time) error {
	if to.IsZero() {
		to = time.Now()
	}

	tracks := []*Track{}
	page := 1

	for {
		res, err := l.recentTracks(ctx, page, from.Unix(), to.Unix())
		if err != nil {
			return err
		}

		if res.Tracks == nil {
			return fmt.Errorf("response tracks is nil")
		}

		for _, rawTrack := range res.Tracks {
			if rawTrack.NowPlaying {
				continue
			}

			tracks = append(tracks, rawTrack.convert())
		}

		if res.Page < res.TotalPages {
			page++
		} else {
			break
		}
	}

	return l.saveHistory(tracks)
}

func (l *lastFm) saveHistory(history []*Track) error {
	if len(history) == 0 {
		return nil
	}

	h := map[int]map[time.Month][]*Track{}

	for i := range history {
		year := history[i].Date.UTC().Year()
		month := history[i].Date.UTC().Month()

		if _, ok := h[year]; !ok {
			h[year] = map[time.Month][]*Track{}
		}

		if _, ok := h[year][month]; !ok {
			h[year][month] = []*Track{}
		}

		h[year][month] = append(h[year][month], history[i])
	}

	for year := range h {
		for month := range h[year] {
			bytes, err := json.MarshalIndent(h[year][month], "", "\t")
			if err != nil {
				return err
			}

			err = os.WriteFile(filepath.Join(l.Output, fmt.Sprintf("%04d-%02d.json", year, month)), bytes, 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *lastFm) recentTracks(ctx context.Context, page int, from, to int64) (*tracks, error) {
	limit := 200
	u, err := url.Parse("https://ws.audioscrobbler.com/2.0/")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("method", "user.getrecenttracks")
	q.Set("user", l.User)
	q.Set("limit", strconv.Itoa(limit))
	q.Set("page", strconv.Itoa(page))
	q.Set("api_key", l.Key)

	if from != 0 {
		q.Set("from", strconv.FormatInt(from, 10))
	}

	if to != 0 {
		q.Set("to", strconv.FormatInt(to, 10))
	}

	u.RawQuery = q.Encode()

	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response *recentTracksResponse
	err = xml.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.RecentTracks, nil
}
