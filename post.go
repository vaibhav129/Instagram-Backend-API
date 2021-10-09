package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"path"
	"time"
)

func createpost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var post Posts
	_ = json.NewDecoder(request.Body).Decode(&post)
	collection := client.Database("appointy").Collection("app")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, post)
	json.NewEncoder(response).Encode(result)

}
func CheckpostwithID(id primitive.ObjectID) (Users, error) {
	var pt Posts
	collection := client.Database("appointy").Collection("meetings")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Posts{ID: id}).Decode(&pt)
	if pt.ID != id {
		err = errors.New("Error 400: ID not present")
	}
	return pt, err
}
func GetPostwithID(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
		return
	}
	response.Header().Set("content-type", "application/json")
	fmt.Println(path.Base(request.URL.Path))
	id, _ := primitive.ObjectIDFromHex(path.Base(request.URL.Path))
	PostID, err := CheckpostwithID(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(PostID)

}
