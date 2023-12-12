package xkom

import (
	"bytes"
	"encoding/json"
	"fmt"
	ihttp "xkomopener/internal/http"
	"xkomopener/internal/utils/helpers"

	http "github.com/vimbing/fhttp"
)

func (s *Scraper) RegisterToNewsletter() error {
	return helpers.GetDefaultRetry().Try(func() error {
		fmt.Println("Registering to newsletter...")

		payload, err := json.Marshal(RegisterToNewsletterPayload{
			ConsentValues: []ConsentValues{
				{
					IsSelected: true,
					ViewStatus: "Viewed",
					Code:       "email_contact",
				},
			},
			ConsentOrigin: "nw_xkom_unbox",
		})

		if err != nil {
			return err
		}

		res, err := s.internal.Client.Request(&ihttp.Request{
			Url:    "https://mobileapi.x-kom.pl/api/v1/xkom/Account/Consents?",
			Method: "PUT",
			Headers: http.Header{
				"host":            {"mobileapi.x-kom.pl"},
				"accept":          {"application/json"},
				"time-zone":       {"UTC"},
				"Authorization":   {fmt.Sprintf("Bearer %s", s.internal.AccountTokens.AccessToken)},
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

		fmt.Println("Successfully registered to newsletter!")

		return nil
	})
}
