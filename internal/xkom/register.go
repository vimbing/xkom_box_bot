package xkom

import (
	"bytes"
	"encoding/json"
	"fmt"
	ihttp "xkomopener/internal/http"
	"xkomopener/internal/utils/helpers"

	http "github.com/vimbing/fhttp"
)

func (s *Scraper) Register() error {
	return helpers.GetDefaultRetry().Try(func() error {
		fmt.Printf("Creating account for: %s...\n", s.internal.AccountData.Email)

		payload, err := json.Marshal(RegisterPayload{
			AccountAuthData: AccountAuthData{
				Password: s.internal.AccountData.Password,
				Username: s.internal.AccountData.Email,
			},
			Consents: []Consents{
				{
					Code:       "regulations",
					IsSelected: true,
				},
				{
					Code:       "offer_adaptin",
					IsSelected: true,
				},
				{
					Code:       "email_contact",
					IsSelected: true,
				},
			},
			AccountIdentity: AccountIdentity{
				Email:     s.internal.AccountData.Email,
				FirstName: s.internal.AccountData.FirstName,
				LastName:  s.internal.AccountData.LastName,
			},
		})

		if err != nil {
			return err
		}

		res, err := s.internal.Client.Request(&ihttp.Request{
			Url:    "https://mobileapi.x-kom.pl/api/v1/xkom/Account",
			Method: "POST",
			Headers: http.Header{
				"host":            {"mobileapi.x-kom.pl"},
				"accept":          {"application/json"},
				"time-zone":       {"UTC"},
				"clientversion":   {"1.85.0"},
				"x-api-key":       {"ushoh9OoY7eerae8aiGh"},
				"accept-language": {"pl-PL;q=1, en-GB;q=0.9"},
				"user-agent":      {"x-kom_prod/1.85.0 (iPhone; iOS 16.0; Scale/2.00)"},
				"content-type":    {"application/json"},
			},
			Body: bytes.NewBuffer(payload),
		})

		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			return ErrStatusCode
		}

		fmt.Println("Account successfully created!")

		return nil
	})
}
