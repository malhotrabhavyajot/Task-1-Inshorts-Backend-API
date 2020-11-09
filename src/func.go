package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Articles struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Subtitle  string             `json:"subtitle,omitempty" bson:"subtitle,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	Timestamp time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	filter := bson.M{}
	var limit int64 = 10
	var page int64 = 1
	collection := client.Database("theappointy").Collection("article")
	projection := bson.D{
		{"title", "Joe"},
		{"subtitle", "Gold"},
	}

	paginatedData, err := New(collection).Limit(limit).Page(page).Filter(filter).Find()
	if err != nil {
		panic(err)
	}
	var lists []Articles
	for _, raw := range paginatedData.Data {
		var article *Articles
		if marshallErr := bson.Unmarshal(raw, &article); marshallErr == nil {
			lists = append(lists, *article)
		}

	}
	fmt.Printf("Norm Find Data: %+v\n", lists)
	fmt.Printf("Normal find pagination info: %+v\n", paginatedData.Pagination)
}
