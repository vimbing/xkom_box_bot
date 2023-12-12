package imap

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"xkomopener/internal/db"
	ihttp "xkomopener/internal/http"
	"xkomopener/internal/proxy"
	"xkomopener/internal/xkom"

	"github.com/BrianLeishman/go-imap"
	"github.com/percolate/retry"
	"github.com/vimbing/cclient"
	http "github.com/vimbing/fhttp"
	tls "github.com/vimbing/utls"
	"golang.org/x/exp/slices"
)

func (c *Client) getUrlFromBody(body string, firstRegexp *regexp.Regexp) (string, error) {
	secondRegexp := regexp.MustCompile(`https:\/\/click\.mail\.x-kom\.pl\/\?qs=[^"]+`)

	url := secondRegexp.FindString(
		firstRegexp.FindString(body),
	)

	if len(url) < 10 {
		return url, errors.New("short_url")
	}

	return url, nil
}

func (c *Client) GetConfirmationLinks() error {
	imap.Verbose = false
	imap.RetryCount = 3

	im, err := imap.New(
		c.Email,
		c.Password,
		"imap.gmail.com",
		993,
	)

	if err != nil {
		return err
	}

	defer im.Clone()

	err = im.SelectFolder("xkom")

	if err != nil {
		return err
	}

	uids, err := im.GetUIDs("ALL")

	if err != nil {
		return err
	}

	slices.Reverse(uids)

	var consideredUids []int

	if len(uids) <= 500 {
		consideredUids = uids
	} else {
		consideredUids = uids[:500]
	}

	emails, err := im.GetEmails(consideredUids...)

	if err != nil {
		return err
	}

	if len(emails) == 0 {
		return nil
	}

	for _, email := range emails {
		if email == nil {
			continue
		}

		var confirmUrlRegexp *regexp.Regexp

		switch {
		case strings.Contains(email.Subject, "zapis i korzystaj"):
			confirmUrlRegexp = regexp.MustCompile(`https:\/\/click\.mail\.x-kom\.pl\/\?qs=\w+.+dostawać newsletter`)
		case strings.Contains(email.Subject, "newsletter i otrzymaj bon na zakupy") || strings.Contains(email.Subject, "zapis do newslettera"):
			confirmUrlRegexp = regexp.MustCompile(`https:\/\/click\.mail\.x-kom\.pl\/\?qs=\w+.+aktywuj newsletter`)
		default:
			continue
		}

		var usedAccount *xkom.AccountData

		for _, acc := range c.ConsideredAccounts {
			if strings.EqualFold(acc.Email, email.To.String()) {
				usedAccount = acc
				break
			}
		}

		if usedAccount == nil {
			continue
		}

		confirmUrl, err := c.getUrlFromBody(email.HTML, confirmUrlRegexp)

		if err != nil {
			continue
		}

		usedAccount.ConfirmationUrl = confirmUrl
	}

	return nil
}

func (c *Client) getNextConfirmationChainStep(entryUrl string) (string, error) {
	nextUrl := ""

	return nextUrl, retry.Re{Max: 5, Delay: time.Millisecond * 100}.Try(func() error {
		client, err := cclient.NewClient(tls.HelloIOS_12_1, proxy.GlobalProxyStorage.GetRandomProxy(), false, 20)

		if err != nil {
			return err
		}

		req, err := http.NewRequest("GET", entryUrl, nil)

		if err != nil {
			return err
		}

		res, err := client.Do(req)

		if err != nil {
			return err
		}

		res.Body.Close()

		if len(res.Header["Location"]) == 0 {
			return nil
		}

		nextUrl = res.Header["Location"][0]

		return nil
	})
}

func (c *Client) extractSalesForceKey(url string) string {
	return regexp.MustCompile(`\w+`).FindString(
		regexp.MustCompile(`\/\w+`).FindString(
			regexp.MustCompile(`zapis\/\w+`).FindString(url),
		),
	)
}

func (c *Client) confirmNewsletter(key string) error {
	return retry.Re{Max: 2, Delay: time.Millisecond * 100}.Try(func() error {
		fmt.Println("Confirming newsletter...")

		body, err := json.Marshal(map[string]string{
			"salesforceKey": key,
		})

		if err != nil {
			return err
		}

		res, err := c.internal.Client.Request(&ihttp.Request{
			Url:    "https://mobileapi.x-kom.pl/api/v1/xkom/newsletter/confirm",
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
			Body: bytes.NewBuffer(body),
		})

		if err != nil {
			return err
		}

		if res.StatusCode != 200 {
			return errors.New("status")
		}

		defer res.Body.Close()

		return nil
	})
}

func (c *Client) getSalesForceKey(entryUrl string) (string, error) {
	var err error
	currentUrl := entryUrl

	for i := 0; i <= 6; i++ {
		if i == 5 {
			return "", errors.New("not_found")
		}

		currentUrl, err = c.getNextConfirmationChainStep(currentUrl)

		if err != nil {
			continue
		}

		if strings.Contains(currentUrl, "podziekowanie") {
			break
		}
	}

	salesForceKey := c.extractSalesForceKey(currentUrl)

	if len(salesForceKey) == 0 {
		return "", errors.New("not_found")
	}

	return salesForceKey, nil
}

func (c *Client) BulkConfirmLinks(accounts []*xkom.AccountData) {
	threads := 8
	workingThreads := 0

	for _, acc := range accounts {
		for {
			if workingThreads < threads {
				break
			}

			time.Sleep(time.Millisecond * 500)
		}

		workingThreads++
		go func(acc *xkom.AccountData) {
			defer func() {
				workingThreads--
			}()

			salesForceKey, err := c.getSalesForceKey(acc.ConfirmationUrl)

			if err != nil {
				return
			}

			err = c.confirmNewsletter(salesForceKey)

			if err != nil {
				return
			}

			fmt.Printf("Confirmed for: %s\n", acc.Email)

			db.AccountsCollection().InsertOne(
				context.TODO(),
				db.Account{
					Email:        acc.Email,
					Password:     acc.Password,
					LastOpenedAt: time.Now(),
				},
			)
		}(acc)
	}
}
