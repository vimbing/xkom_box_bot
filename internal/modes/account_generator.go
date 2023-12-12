package modes

import (
	"fmt"
	"os"
	"time"
	"xkomopener/internal/imap"
	"xkomopener/internal/xkom"
)

func AccountGenerator() {
	accounts := []*xkom.AccountData{}

	count := 65

	threads := 8
	workingThreads := 0

	for {
		if len(accounts) >= count {
			break
		}

		if workingThreads >= threads {
			time.Sleep(time.Millisecond * 200)
			continue
		}

		workingThreads++

		go func() {
			fmt.Printf("%d / %d\n", len(accounts), count)

			defer func() {
				time.Sleep(time.Millisecond * 250)
				workingThreads--
			}()

			xkomScraper, err := xkom.Init()

			if err != nil {
				return
			}

			account, err := xkomScraper.GenerateAccount()

			if err != nil {
				return
			}

			if account == nil {
				return
			}

			accounts = append(accounts, account)
		}()
	}

	sleepSeconds := 60 * 4
	for i := 0; i < sleepSeconds; i++ {
		fmt.Printf("Sleeping for %ds\n", sleepSeconds-i)
		time.Sleep(time.Second)
	}

	fmt.Println("Handling imap...")

	imapClient, err := imap.Init(os.Getenv("IMAP_EMAIL"), os.Getenv("IMAP_PASSWORD"), accounts)

	if err != nil {
		return
	}

	err = imapClient.GetConfirmationLinks()

	if err != nil {
		return
	}

	imapClient.BulkConfirmLinks(accounts)
}
