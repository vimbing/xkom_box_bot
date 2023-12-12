package webhook

import "xkomopener/internal/http"

func Init(url string) (*Webhook, error) {
	client, err := http.Init()

	if err != nil {
		return &Webhook{}, err
	}

	return &Webhook{
		Url:     url,
		Payload: &WebhookPayload{},
		client:  client,
	}, nil
}
