package db

import "time"

type Account struct {
	Email        string    `bson:"email"  json:"email"`
	Password     string    `bson:"password"  json:"password"`
	LastOpenedAt time.Time `bson:"lastOpenedAt"  json:"lastOpenedAt"`
}

type Promotion struct {
	Email string `bson:"email" json:"email"`
	Grade string `bson:"grade" json:"grade"`
}
