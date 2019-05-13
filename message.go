package traq

import (
	"fmt"
	"net/http"
)

type PostMessageResponse struct {
	MessageResponse
}

// PostMessage チャンネルにメッセージを投稿します
//
// POST /api/1.0/channels/{channelID}/messages
func (c *Client) PostMessage(channelID, text string) (*PostMessageResponse, error) {
	var (
		success PostMessageResponse
		failure ErrorResponse
	)

	path := fmt.Sprintf("%s/channels/%s/messages", v1Prefix, channelID)
	resp, err := c.sling.New().
		Post(path).
		BodyJSON(map[string]string{"text": text}).
		Receive(&success, &failure)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		failure.Code = resp.StatusCode
		return nil, &failure
	}
	return &success, nil
}

// PostMessage ユーザーにダイレクトメッセージを送信します
//
// POST /api/1.0/users/{userID}/messages
func (c *Client) PostDirectMessage(userID, text string) (*PostMessageResponse, error) {
	var (
		success PostMessageResponse
		failure ErrorResponse
	)

	path := fmt.Sprintf("%s/users/%s/messages", v1Prefix, userID)
	resp, err := c.sling.New().
		Post(path).
		BodyJSON(map[string]string{"text": text}).
		Receive(&success, &failure)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		failure.Code = resp.StatusCode
		return nil, &failure
	}
	return &success, nil
}
