package traq

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type PostMessageViaWebhookOpt struct {
	Secret  string
	BaseURL string
}

// PostMessageViaWebhook Webhookメッセージを送信します
//
// POST /api/1.0/webhooks/{webhookID}
func PostMessageViaWebhook(webhookID, text string, opt *PostMessageViaWebhookOpt) error {
	var (
		failure ErrorResponse
	)

	base := DefaultBaseURL
	sig := ""
	if opt != nil {
		if len(opt.BaseURL) > 0 {
			base = opt.BaseURL
		}
		if len(opt.Secret) > 0 {
			mac := hmac.New(sha1.New, []byte(opt.Secret))
			_, _ = mac.Write([]byte(text))
			sig = hex.EncodeToString(mac.Sum(nil))
		}
	}

	path := fmt.Sprintf("%s/webhooks/%s", v1Prefix, webhookID)
	resp, err := sling.New().
		Base(base).
		Post(path).
		BodyProvider(&plainTextProvider{text: text}).
		Set("X-TRAQ-Signature", sig).
		Receive(nil, &failure)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		failure.Code = resp.StatusCode
		return &failure
	}
	return nil
}
