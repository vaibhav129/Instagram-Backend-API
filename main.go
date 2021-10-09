package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	fmt.Println("Application Running")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, _ = mongo.Connect(ctx, clientOptions)

	http.HandleFunc("/users", createuser)
	http.HandleFunc("/userid/", GetUserwithID)
	http.HandleFunc("/post", createpost)
	http.HandleFunc("/postid/", GetPostwithID)
	http.ListenAndServe(":12345", nil)
}


