package imap

import (
	"xkomopener/internal/http"
	"xkomopener/internal/xkom"
)

func Init(email, password string, consideredAccounts []*xkom.AccountData) (*Client, error) {
	httpClient, err := http.Init()

	if err != nil {
		return &Client{}, err
	}

	return &Client{
		internal: internal{
			Client: httpClient,
		},
		Email:              email,
		Password:           password,
		ConsideredAccounts: consideredAccounts,
	}, nil
}
