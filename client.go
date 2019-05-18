package traq

import (
	"context"
	"github.com/dghubble/sling"
	"golang.org/x/oauth2"
)

const (
	DefaultBaseURL = "https://q.trap.jp"
	v1Prefix       = "/api/1.0"
)

type Client struct {
	baseURL string
	sling   *sling.Sling
}

func NewClient(base string, accessToken string) *Client {
	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: accessToken}
	httpClient := config.Client(context.Background(), token)

	return &Client{
		baseURL: base,
		sling:   sling.New().Client(httpClient).Base(base),
	}
}
