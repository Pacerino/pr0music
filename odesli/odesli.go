package odesli

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type odesliResponse struct {
	EntityUniqueID  string `json:"entityUniqueId"`
	UserCountry     string `json:"userCountry"`
	PageURL         string `json:"pageUrl"`
	LinksByPlatform struct {
		Spotify struct {
			Country             string `json:"country"`
			URL                 string `json:"url"`
			NativeAppURIDesktop string `json:"nativeAppUriDesktop"`
			EntityUniqueID      string `json:"entityUniqueId"`
		} `json:"spotify"`
		Youtube struct {
			Country        string `json:"country"`
			URL            string `json:"url"`
			EntityUniqueID string `json:"entityUniqueId"`
		} `json:"youtube"`
	} `json:"linksByPlatform"`
}

type rlhttpclient struct {
	client      *http.Client
	Ratelimiter *rate.Limiter
}

func (c *rlhttpclient) Do(req *http.Request) (*http.Response, error) {
	ctx := context.Background()
	err := c.Ratelimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewClient(rl *rate.Limiter) *rlhttpclient {
	c := &rlhttpclient{
		client:      http.DefaultClient,
		Ratelimiter: rl,
	}
	return c
}

func GetLinks(url string) (*odesliResponse, error) {
	rl := rate.NewLimiter(rate.Every(60*time.Second), 10) // 10 request every 1 minute
	c := NewClient(rl)
	req, _ := http.NewRequest("GET", "https://api.song.link/v1-alpha.1/links?songIfSingle=true&platform=spotify,youtube&type=song&url="+url, nil)
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var odesliResp odesliResponse
	err = json.NewDecoder(resp.Body).Decode(&odesliResp)
	if err != nil {
		return nil, err
	}

	return &odesliResp, nil
}
