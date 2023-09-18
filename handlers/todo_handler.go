package todo_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aboobakersiddiqr63/go-crud/helper"
	todo "github.com/aboobakersiddiqr63/go-crud/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	LoadEnv()
	db_uri := os.Getenv("DB_URI")
	db_name := os.Getenv("DB_NAME")
	db_collection := os.Getenv("DB_COLLECTION_NAME")

	clientOption := options.Client().ApplyURI(db_uri)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(db_name).Collection(db_collection)

	fmt.Println("Collection instance is ready")
}

func LoadEnv() {
	err := godotenv.Load(".env")
	helper.HandleException(err, "LoadEnv")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	response := getAllTasks()
	json.NewEncoder(w).Encode(response)
}

func Createtask(w http.ResponseWriter, r *http.Request) {
	helper.SetCommonHeaders(w, "POST")

	//decode th request body
	var task todo.ToDoList
	json.NewDecoder(r.Body).Decode(&task)

	response := createtask(task)
	json.NewEncoder(w).Encode(response)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	helper.SetCommonHeaders(w, "PUT")

	params := mux.Vars(r)

	response := taskComplete(params["id"])
	json.NewEncoder(w).Encode(response)
}

func UndoTaskStatus(w http.ResponseWriter, r *http.Request) {
	helper.SetCommonHeaders(w, "PUT")

	params := mux.Vars(r)

	response := undoTaskStatus(params["id"])
	json.NewEncoder(w).Encode(response)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	helper.SetCommonHeaders(w, "DELETE")

	params := mux.Vars(r)

	response := deleteTask(params["id"])
	json.NewEncoder(w).Encode(response)
}

func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	helper.SetCommonHeaders(w, "DELETE")

	response := deleteAllTask()
	json.NewEncoder(w).Encode(response)
}

//

func createtask(task todo.ToDoList) string {
	_, err := collection.InsertOne(context.Background(), task)
	helper.HandleException(err, "createtask")
	response := fmt.Sprintln("Task is successfully created")
	return response
}

func getAllTasks() []primitive.M {
	fmt.Println("debug error four")

	result, err := collection.Find(context.Background(), bson.D{{}})
	fmt.Println("debug error five")

	helper.HandleException(err, "getAllTasks")

	var response []primitive.M
	for result.Next(context.Background()) {
		var resp bson.M
		fmt.Println("Method reached here")
		e := result.Decode(&resp)
		if e != nil {
			log.Fatal(e)
		}
		response = append(response, resp)
	}

	defer result.Close(context.Background())
	return response
}

func taskComplete(taskId string) string {
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	helper.HandleException(err, "taskComplete")

	response := fmt.Sprintln("Task is marked as completed")
	return response
}

func undoTaskStatus(taskId string) string {
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	helper.HandleException(err, "undoTaskStatus")

	response := fmt.Sprintln("Task is marked as incomplete")
	return response
}

func deleteTask(taskId string) string {
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	helper.HandleException(err, "deleteTask")

	response := fmt.Sprintln("Task is deleted")
	return response
}

func deleteAllTask() string {
	_, err := collection.DeleteMany(context.Background(), bson.M{})
	helper.HandleException(err, "deleteAllTask")
	response := fmt.Sprintln("All the tasks removed")
	return response
}
