package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-manager/internal/handlers"
	"task-manager/internal/models"
	"task-manager/internal/services"
	"task-manager/internal/storage"
	"testing"
)

func setupTest() *handlers.TaskHandler {
	taskStorage := storage.NewMemoryStorage()
	taskService := services.NewTaskService(taskStorage)
	return handlers.NewTaskHandler(*taskService)
}

func TestHandleTasks_POST_GET(t *testing.T) {
	handler := setupTest()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", handler.HandleTasks)
	mux.HandleFunc("POST /tasks", handler.HandleTasks)
	server := httptest.NewServer(mux)
	defer server.Close()

	// Test POST
	task := map[string]string{"title": "Test task", "description": "Test description"}
	jsonData, _ := json.Marshal(task)

	resp, err := http.Post(server.URL+"/tasks", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Could not send POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	// Test GET
	resp, err = http.Get(server.URL + "/tasks")
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var tasks []models.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestHandleTasks_InvalidMethods(t *testing.T) {
	handler := setupTest()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", handler.HandleTasks)
	mux.HandleFunc("POST /tasks", handler.HandleTasks)
	server := httptest.NewServer(mux)
	defer server.Close()

	// Test PUT (не должен работать)
	req, err := http.NewRequest("PUT", server.URL+"/tasks", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405 for PUT, got %d", resp.StatusCode)
	}
}

func TestHandleTaskByID_InvalidMethods(t *testing.T) {
	handler := setupTest()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks/{id}", handler.HandleTaskByID)
	mux.HandleFunc("DELETE /tasks/{id}", handler.HandleTaskByID)
	server := httptest.NewServer(mux)
	defer server.Close()

	// Test POST (not allowed)
	req, err := http.NewRequest("POST", server.URL+"/tasks/1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405 for POST, got %d", resp.StatusCode)
	}
}
