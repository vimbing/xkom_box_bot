package imap

import (
	"xkomopener/internal/http"
	"xkomopener/internal/xkom"
)

type internal struct {
	Client *http.Client
}

type Client struct {
	internal           internal
	Email              string
	Password           string
	ConsideredAccounts []*xkom.AccountData
}
