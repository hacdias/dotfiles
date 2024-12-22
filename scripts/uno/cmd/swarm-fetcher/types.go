package main

import (
	"sort"
	"time"
)

type swarmMeta struct {
	Code int `json:"code"`
}

type swarmCheckin map[string]interface{}

func (c swarmCheckin) CreatedAt() time.Time {
	v, ok := c["createdAt"]
	if !ok {
		panic("createdAt not available")
	}

	i, ok := v.(float64)
	if !ok {
		panic("createdAt is not an float64")
	}

	return time.Unix(int64(i), 0)
}

type swarmCheckins []swarmCheckin

func (c swarmCheckins) sort() {
	sort.SliceStable(c, func(i, j int) bool {
		return c[i].CreatedAt().After(c[j].CreatedAt())
	})
}

type swarmResponse struct {
	Checkins struct {
		Items swarmCheckins `json:"items"`
	} `json:"checkins"`
}

type swarmOutput struct {
	Meta     swarmMeta     `json:"meta"`
	Response swarmResponse `json:"response"`
}
