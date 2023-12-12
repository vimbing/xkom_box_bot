package migrations

import (
	"context"
	"time"
	"xkomopener/internal/db"

	"go.mongodb.org/mongo-driver/bson"
)

func MigrateDates() {
	db.AccountsCollection().UpdateMany(
		context.TODO(),
		bson.M{},
		bson.M{
			"$set": bson.M{
				"lastOpenedAt": time.Now(),
			},
		},
	)
}
