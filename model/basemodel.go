package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	fmt.Println("Connecting to mongo")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// var err error
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// defer func() {
	// 	fmt.Println("defer func")
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		fmt.Println("Error disconnecting mongo client")
	// 		panic(err)
	// 	}
	// }()
}
