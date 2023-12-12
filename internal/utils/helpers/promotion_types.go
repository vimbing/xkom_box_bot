package helpers

import (
	"context"
	"fmt"
	"os"
	"xkomopener/internal/db"

	"go.mongodb.org/mongo-driver/bson"
)

func ShowTypesOfPromotionsInDb() {
	cursor, err := db.PromotionsCollection().Find(
		context.TODO(),
		bson.M{},
	)

	if err != nil {
		panic(err)
	}

	promotions := db.UnmarshalFind[db.Promotion](cursor)

	count := make(map[string]int)

	for _, promotion := range promotions {
		if _, ok := count[promotion.Grade]; !ok {
			count[promotion.Grade] = 0
			continue
		}

		count[promotion.Grade]++
	}

	fmt.Println("Current promotions: ")

	for k, v := range count {
		fmt.Printf("%s: %d\n", k, v)
	}

	os.Exit(0)
}
