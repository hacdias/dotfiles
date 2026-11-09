package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"uno"
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
		res, err := l.recentTracks(ctx, page, from, to)
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

	return uno.SaveByMonth(l.Output, tracks, func(t *Track) time.Time { return t.Date })
}

func (l *lastFm) recentTracks(ctx context.Context, page int, from, to time.Time) (*tracks, error) {
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

	if !from.IsZero() {
		q.Set("from", strconv.FormatInt(from.Unix(), 10))
	}

	if !to.IsZero() {
		q.Set("to", strconv.FormatInt(to.Unix(), 10))
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

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status is %s", res.Status)
	}

	var response *recentTracksResponse
	err = xml.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.RecentTracks, nil
}
