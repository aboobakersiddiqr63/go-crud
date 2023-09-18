package router

import (
	todo_handler "github.com/aboobakersiddiqr63/go-crud/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", todo_handler.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", todo_handler.Createtask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", todo_handler.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undo/task/{id}", todo_handler.UndoTaskStatus).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task/{id}", todo_handler.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/tasks", todo_handler.DeleteAllTask).Methods("DELETE", "OPTIONS")

	return router
}
