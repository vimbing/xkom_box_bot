package xkom

import (
	"xkomopener/internal/http"
	"xkomopener/internal/utils/config"
)

type AccountTokens struct {
	AccessToken  string
	RefreshToken string
}

type internal struct {
	Client        *http.Client
	AccountData   *AccountData
	BotConfig     *config.BotConfig
	AccountTokens AccountTokens
}

type Scraper struct {
	internal internal
}

type AccountData struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	ConfirmationUrl string `json:"confirmationUrl"`
}

type BoxType int

type BoxOpenResult struct {
	Img            string
	RarityName     string
	ItemName       string
	Price          float64
	PromotionValue float64
}
