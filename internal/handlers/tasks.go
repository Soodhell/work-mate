package handlers

import (
	"encoding/json"
	"net/http"
	"task-manager/internal/services"
)

type TaskHandler struct {
	service services.TaskService
}

func NewTaskHandler(service services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createTask(w, r)
	case http.MethodGet:
		h.getTasks(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getTask(w, r, id)
	case http.MethodDelete:
		h.deleteTask(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	task := h.service.CreateTask()
	respondWithJSON(w, http.StatusCreated, task)
}

func (h *TaskHandler) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.service.GetAllTasks()
	respondWithJSON(w, http.StatusOK, tasks)
}

func (h *TaskHandler) getTask(w http.ResponseWriter, r *http.Request, taskID string) {
	task, exists := h.service.GetTask(taskID)
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	respondWithJSON(w, http.StatusOK, task)
}

func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request, taskID string) {
	h.service.DeleteTask(taskID)
	w.WriteHeader(http.StatusNoContent)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
