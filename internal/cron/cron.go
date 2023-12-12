package cron

import (
	"xkomopener/internal/modes"

	"github.com/robfig/cron/v3"
)

func Cron() {
	c := cron.New()

	c.AddFunc("0 * * * *", modes.OpenBoxes)

	c.Start()
}
