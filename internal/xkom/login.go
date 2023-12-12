package xkom

import (
	"fmt"
	"net/url"
	"strings"
	ihttp "xkomopener/internal/http"
	"xkomopener/internal/utils/helpers"

	http "github.com/vimbing/fhttp"
)

func (s *Scraper) Login() error {
	return helpers.GetDefaultRetry().Try(func() error {
		fmt.Println("Logging in...")

		payload := fmt.Sprintf(`client_id=&client_secret=&grant_type=password&password=%s&scope=user&username=%s`, url.QueryEscape(s.internal.AccountData.Password), url.QueryEscape(s.internal.AccountData.Email))

		res, err := s.internal.Client.Request(&ihttp.Request{
			Url:    "https://mobileapi.x-kom.pl/api/v1/xkom/Token",
			Method: "POST",
			Headers: http.Header{
				"host":            {"mobileapi.x-kom.pl"},
				"accept":          {"application/json"},
				"time-zone":       {"UTC"},
				"clientversion":   {"1.85.0"},
				"x-api-key":       {"ushoh9OoY7eerae8aiGh"},
				"accept-language": {"pl-PL;q=1, en-GB;q=0.9"},
				"user-agent":      {"x-kom_prod/1.85.0 (iPhone; iOS 16.0; Scale/2.00)"},
				"content-type":    {"application/x-www-form-urlencoded"},
			},
			Body: strings.NewReader(payload),
		})

		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			return ErrStatusCode
		}

		var resJson LoginResponse
		s.internal.Client.Decode(res, &resJson)

		s.internal.AccountTokens = AccountTokens{
			AccessToken:  resJson.AccessToken,
			RefreshToken: resJson.RefreshToken,
		}

		fmt.Println("Logged in successfully!")

		return nil
	})
}
