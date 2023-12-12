package xkom

type RegisterPayload struct {
	AccountAuthData AccountAuthData `json:"AccountAuthData"`
	Consents        []Consents      `json:"Consents"`
	AccountIdentity AccountIdentity `json:"AccountIdentity"`
}

type AccountAuthData struct {
	Password string `json:"Password"`
	Username string `json:"Username"`
}

type Consents struct {
	Code       string `json:"Code"`
	IsSelected bool   `json:"IsSelected"`
}

type AccountIdentity struct {
	Email     string `json:"Email"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type RegisterToNewsletterPayload struct {
	ConsentValues []ConsentValues `json:"consentValues"`
	ConsentOrigin string          `json:"consentOrigin"`
}

type ConsentValues struct {
	IsSelected bool   `json:"isSelected"`
	ViewStatus string `json:"viewStatus"`
	Code       string `json:"code"`
}
