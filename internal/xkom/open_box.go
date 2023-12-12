package xkom

import (
	"fmt"
	ihttp "xkomopener/internal/http"
	"xkomopener/internal/utils/helpers"

	http "github.com/vimbing/fhttp"
)

func (s *Scraper) OpenBox(boxType BoxType) (BoxOpenResult, error) {
	var openResult BoxOpenResult

	return openResult, helpers.GetDefaultRetry().Try(func() error {
		fmt.Printf("Opening box with type: %d\n", boxType)

		fmt.Println(s.internal.AccountData.Email)

		res, err := s.internal.Client.Request(&ihttp.Request{
			Url:    fmt.Sprintf("https://mobileapi.x-kom.pl/api/v1/xkom/Box/%d/Roll?", boxType),
			Method: "POST",
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
		})

		if err != nil {
			return err
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			return ErrStatusCode
		}

		var resJson OpenBoxResponse
		s.internal.Client.Decode(res, &resJson)

		openResult = BoxOpenResult{
			RarityName:     resJson.BoxRarity.Name,
			ItemName:       resJson.Item.Name,
			Price:          resJson.BoxPrice,
			PromotionValue: resJson.PromotionGain.Value,
			Img:            resJson.Item.Photo.URL,
		}

		fmt.Println("Box opened successfully!")

		return nil
	})
}
