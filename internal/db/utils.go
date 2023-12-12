package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func UnmarshalFind[ResultType interface{}](cursor *mongo.Cursor) []ResultType {
	resultArr := make([]ResultType, 0)

	for cursor.TryNext(context.TODO()) {
		var result ResultType

		err := cursor.Decode(&result)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		resultArr = append(resultArr, result)
	}

	return resultArr
}
