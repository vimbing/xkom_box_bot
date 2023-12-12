package webhook

import (
	"time"

	"github.com/google/uuid"
)

func (q *Queue) StartHandler() {
	for {
		time.Sleep(time.Millisecond * 100)

		if len(q.Webhooks) == 0 {
			continue
		}

		q.Mutex.Lock()

		for k, w := range q.Webhooks {
			w.Send()
			time.Sleep(time.Millisecond * 500)
			delete(q.Webhooks, k)
		}

		q.Mutex.Unlock()
	}
}

func InitWebhookQueue() {
	go webhookQueue.StartHandler()
}

func (q *Queue) AddToQueue(w *Webhook) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()

	q.Webhooks[uuid.NewString()] = w
}

func AddWebhookToQueue(w *Webhook) {
	webhookQueue.AddToQueue(w)
}
