package db

import "go.mongodb.org/mongo-driver/mongo"

func GetDatabase() *mongo.Database {
	return DbClient.Database(DB_NAME)
}

func AccountsCollection() *mongo.Collection {
	return GetDatabase().Collection(ACCOUNTS_COLLECTION)
}

func PromotionsCollection() *mongo.Collection {
	return GetDatabase().Collection(PROMOTIONS_COLLECTION)
}
