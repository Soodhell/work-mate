package main

import (
	"log"
	"net/http"
	"task-manager/internal/handlers"
	"task-manager/internal/services"
	"task-manager/internal/storage"
)

func main() {
	taskStorage := storage.NewMemoryStorage()

	taskService := services.NewTaskService(taskStorage)

	taskHandler := handlers.NewTaskHandler(*taskService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", taskHandler.HandleTasks)
	mux.HandleFunc("POST /tasks", taskHandler.HandleTasks)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.HandleTaskByID)
	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.HandleTaskByID)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
