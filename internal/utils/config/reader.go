package config

import (
	"encoding/json"
	"os"
)

func ReadBotConfig() (BotConfig, error) {
	fileData, err := os.ReadFile(os.Getenv("BOT_CONFIG_PATH"))

	if err != nil {
		return BotConfig{}, err
	}

	var botConfig BotConfig
	json.Unmarshal(fileData, &botConfig)

	return botConfig, nil
}
