package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Articles struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Subtitle  string             `json:"subtitle,omitempty" bson:"subtitle,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	Timestamp time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

func CreateArticles(r http.ResponseWriter, req *http.Request) {
	r.Header().Set("content-type", "application/json")
	var articles Articles
	_ = json.NewDecoder(req.Body).Decode(&articles)
	collection := client.Database("theappointy").Collection("article")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, articles)
	json.NewEncoder(r).Encode(result)
}

func GetArticlesbyID(r http.ResponseWriter, req *http.Request) {
	r.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var articles Articles
	collection := client.Database("theappointy").Collection("article")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Articles{ID: id}).Decode(&articles)
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(r).Encode(articles)
}

func GetAllArticle(r http.ResponseWriter, req *http.Request) {
	r.Header().Set("content-type", "application/json")
	var article []Articles
	collection := client.Database("theappointy").Collection("article")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var articles Articles
		cursor.Decode(&articles)
		article = append(article, articles)
	}
	if err := cursor.Err(); err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(r).Encode(article)
}

func GetSearchArticles(r http.ResponseWriter, req *http.Request) {
	r.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	search, _ := (params["search"])
	var articles Articles
	collection := client.Database("theappointy").Collection("article")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	arr := collection.FindOne(ctx, Articles{Title: search}).Decode(&articles)
	brr := collection.FindOne(ctx, Articles{Subtitle: search}).Decode(&articles)
	crr := collection.FindOne(ctx, Articles{Content: search}).Decode(&articles)
	if arr != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{ "message": "` + arr.Error() + `" }`))
		return
	}
	if arr != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{ "message": "` + brr.Error() + `" }`))
		return
	}
	if arr != nil {
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte(`{ "message": "` + crr.Error() + `" }`))
		return
	}
	json.NewEncoder(r).Encode(articles)
}
func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/articles", CreateArticles).Methods("POST")
	router.HandleFunc("/articles", GetAllArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", GetArticlesbyID).Methods("GET")
	//router.HandleFunc("/articles", GetSearchArticles).Methods("GET")
	http.ListenAndServe(":12345", router)

}
