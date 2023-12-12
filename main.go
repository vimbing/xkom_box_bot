package main

import (
	"fmt"
	"os"
	"sync"
	"xkomopener/internal/cron"
	"xkomopener/internal/db"
	"xkomopener/internal/modes"
	"xkomopener/internal/proxy"
	"xkomopener/internal/webhook"

	"github.com/joho/godotenv"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	wg.Add(1)

	godotenv.Load()

	fmt.Println(os.Getenv("PROXIES_PATH"))

	if err := proxy.GlobalProxyStorage.LoadFromFile(); err != nil {
		panic(err)
	}

	if err := db.Connect(); err != nil {
		panic(err)
	}

	webhook.InitWebhookQueue()

	cron.Cron()

	modes.AccountGenerator()
	modes.OpenBoxes()

	wg.Wait()
}
