package webhook

var (
	webhookQueue = Queue{Webhooks: make(map[string]*Webhook)}
)
