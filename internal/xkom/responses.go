package xkom

import "time"

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type OpenBoxResponse struct {
	BoxItemRolledResourceID string `json:"BoxItemRolledResourceId"`
	Item                    struct {
		ID    string `json:"Id"`
		Name  string `json:"Name"`
		Photo struct {
			URL          string `json:"Url"`
			ThumbnailURL string `json:"ThumbnailUrl"`
			URLTemplate  string `json:"UrlTemplate"`
		} `json:"Photo"`
		CatalogPrice         float64 `json:"CatalogPrice"`
		CategoryNameSingular string  `json:"CategoryNameSingular"`
		Category             struct {
			ID             string `json:"Id"`
			NameSingular   string `json:"NameSingular"`
			NamePlural     string `json:"NamePlural"`
			Photo          any    `json:"Photo"`
			ParentCategory struct {
				ID           string `json:"Id"`
				NameSingular string `json:"NameSingular"`
				NamePlural   string `json:"NamePlural"`
				Photo        struct {
					URL          string `json:"Url"`
					ThumbnailURL string `json:"ThumbnailUrl"`
					URLTemplate  any    `json:"UrlTemplate"`
				} `json:"Photo"`
				ParentCategory any `json:"ParentCategory"`
				ParentGroup    struct {
					ID   string `json:"Id"`
					Name string `json:"Name"`
				} `json:"ParentGroup"`
			} `json:"ParentCategory"`
			ParentGroup struct {
				ID   string `json:"Id"`
				Name string `json:"Name"`
			} `json:"ParentGroup"`
		} `json:"Category"`
		ProducerName string `json:"ProducerName"`
	} `json:"Item"`
	BoxRarity struct {
		ID   string `json:"Id"`
		Name string `json:"Name"`
	} `json:"BoxRarity"`
	BoxPrice            float64   `json:"BoxPrice"`
	WebURL              string    `json:"WebUrl"`
	AvailableToRollDate time.Time `json:"AvailableToRollDate"`
	PromotionGain       struct {
		Value     float64 `json:"Value"`
		GainValue string  `json:"GainValue"`
		GainType  string  `json:"GainType"`
	} `json:"PromotionGain"`
}
