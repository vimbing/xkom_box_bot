package modes

import (
	"context"
	"fmt"
	"os"
	"time"
	"xkomopener/internal/db"
	"xkomopener/internal/webhook"
	"xkomopener/internal/xkom"

	"go.mongodb.org/mongo-driver/bson"
)

func OpenBoxes() {
	cursor, err := db.AccountsCollection().Find(context.TODO(), bson.M{})

	if err != nil {
		return
	}

	accounts := db.UnmarshalFind[db.Account](cursor)

	for i, account := range accounts {
		if !account.LastOpenedAt.Add(time.Hour * 24).Add(time.Minute * 15).After(time.Now()) {
			continue
		}

		if i <= 15 {
			continue
		}

		xkomScraper, err := xkom.Init()

		if err != nil {
			return
		}

		xkomScraper.SetLoginCredentials(
			account.Email,
			account.Password,
		)

		err = xkomScraper.Login()

		if err != nil {
			return
		}

		for _, boxType := range []xkom.BoxType{xkom.GrayBox, xkom.BlueBox, xkom.PurpleBox} {
			go func(boxType xkom.BoxType, account db.Account) {
				openResult, err := xkomScraper.OpenBox(boxType)

				if err != nil {
					return
				}

				w, err := webhook.Init(os.Getenv("DISCORD_WEBHOOK"))

				if err != nil {
					return
				}

				db.PromotionsCollection().InsertOne(context.TODO(), db.Promotion{
					Email: account.Email,
					Grade: openResult.RarityName,
				})

				w.GetWebhookData().
					AddEmbed().
					SetImage(openResult.Img).
					SetTitle("Box opened!").
					SetColor(boxType.GetWebhookColor()).
					AddField(webhook.Field{Name: "Rarity", Value: openResult.RarityName, Inline: true}).
					AddField(webhook.Field{Name: "Item Name", Value: openResult.ItemName, Inline: true}).
					AddField(webhook.Field{Name: "Price", Value: fmt.Sprintf("%.2f PLN", openResult.Price), Inline: true}).
					AddField(webhook.Field{Name: "Discount", Value: fmt.Sprintf("%.2f PLN", openResult.PromotionValue), Inline: true}).
					AddField(webhook.Field{Name: "Account", Value: fmt.Sprintf("%s:%s", account.Email, account.Password), Inline: true})

				webhook.AddWebhookToQueue(w)
			}(boxType, account)
		}

		time.Sleep(time.Second * 2)

		db.AccountsCollection().UpdateOne(
			context.TODO(),
			bson.M{"email": account.Email},
			bson.M{
				"$set": bson.M{
					"lastOpenedAt": time.Now(),
				},
			},
		)
	}
}
