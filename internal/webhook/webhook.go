package webhook

import (
	"bytes"
	"encoding/json"
	ihttp "xkomopener/internal/http"

	http "github.com/vimbing/fhttp"
)

func (w *WebhookPayload) AddEmbed() *Embed {
	embedIndex := len(w.Embeds)

	w.Embeds = append(w.Embeds, Embed{})

	return &w.Embeds[embedIndex]
}

func (w *WebhookPayload) SetContent(content string) *WebhookPayload {
	w.Content = content
	return w
}

func (w *Webhook) GetWebhookData() *WebhookPayload {
	return w.Payload
}

func (w *Webhook) Send() {
	payload, err := json.Marshal(w.Payload)

	if err != nil {
		return
	}

	w.client.Request(&ihttp.Request{
		Method: "POST",
		Url:    w.Url,
		Body:   bytes.NewBuffer(payload),
		Headers: http.Header{
			"content-type": {"application/json"},
		},
	})
}
