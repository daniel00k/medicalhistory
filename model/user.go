package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Fname string             `bson:"name,omitempty"`
	Lname string             `bson:"lastname,omitempty"`
	Email string             `bson:"email,omitempty"`
}

func (u User) All() []User {
	collection := client.Database("historialmedico").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Printf("asahdjkasdhjkashdkjsa" + err.Error())
	}
	var users []User
	res.All(ctx, &users)
	return users
}

func (u User) Create() User {
	collection := client.Database("historialmedico").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, u)
	if err != nil {
		fmt.Printf("asahdjkasdhjkashdkjsa" + err.Error())
	}

	return User{
		ID:    res.InsertedID.(primitive.ObjectID),
		Fname: u.Fname,
		Lname: u.Lname,
		Email: u.Email,
	}
}
