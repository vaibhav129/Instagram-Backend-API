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

func createuser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Users
	_ = json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("appointy").Collection("app")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)

}
func CheckUserwithID(id primitive.ObjectID) (Users, error) {
	var meet Users
	collection := client.Database("appointy").Collection("app")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Users{ID: id}).Decode(&meet)
	if meet.ID != id {
		err = errors.New("Error 400: ID not present")
	}
	return meet, err
}
func GetUserwithID(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
		return
	}
	response.Header().Set("content-type", "application/json")
	fmt.Println(path.Base(request.URL.Path))
	id, _ := primitive.ObjectIDFromHex(path.Base(request.URL.Path))
	meetingwithID, err := CheckUserwithID(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(meetingwithID)

}
