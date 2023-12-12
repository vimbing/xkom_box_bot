package xkom

import (
	"xkomopener/internal/http"
	"xkomopener/internal/utils/config"
)

func Init() (*Scraper, error) {
	client, err := http.Init()

	if err != nil {
		return &Scraper{}, err
	}

	botConfig, err := config.ReadBotConfig()

	if err != nil {
		return &Scraper{}, err
	}

	return &Scraper{
		internal: internal{
			Client:      client,
			BotConfig:   &botConfig,
			AccountData: &AccountData{},
		},
	}, nil
}
