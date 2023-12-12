package webhook

import (
	"sync"
	"xkomopener/internal/http"
)

type Webhook struct {
	Payload *WebhookPayload
	Url     string

	client *http.Client
}

type WebhookPayload struct {
	Content     any     `json:"content"`
	Embeds      []Embed `json:"embeds"`
	Attachments []any   `json:"attachments"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Embed struct {
	Title  string  `json:"title"`
	Color  int     `json:"color"`
	Fields []Field `json:"fields"`
	Image  Image   `json:"image"`
}

type Image struct {
	Url string `json:"url"`
}

type Queue struct {
	Mutex    sync.RWMutex
	Webhooks map[string]*Webhook
}
