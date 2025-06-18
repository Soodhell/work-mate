package services

import (
	"log"
	"task-manager/internal/models"
	"task-manager/internal/storage"
	"time"
)

type TaskService struct {
	storage storage.TaskStorage
}

func NewTaskService(storage storage.TaskStorage) *TaskService {
	return &TaskService{storage: storage}
}

func (s *TaskService) CreateTask() models.Task {
	task := models.Task{
		ID:        generateID(),
		Status:    models.StatusPending,
		CreatedAt: time.Now(),
	}

	s.storage.Create(task)

	go s.processTask(task.ID)

	return task
}

func (s *TaskService) GetTask(id string) (models.Task, bool) {
	return s.storage.Get(id)
}

func (s *TaskService) DeleteTask(id string) {
	s.storage.Delete(id)
}

func (s *TaskService) processTask(id string) {
	task, exists := s.storage.Get(id)
	if !exists {
		return
	}

	task.Status = models.StatusProcessing
	now := time.Now()
	task.StartedAt = &now
	s.storage.Create(task)

	duration := time.Duration(180+time.Now().UnixNano()%120) * time.Second
	time.Sleep(duration)

	task.Status = models.StatusCompleted
	result := "Task completed successfully after " + duration.String()
	task.Result = &result
	completedAt := time.Now()
	task.CompletedAt = &completedAt
	processingTime := completedAt.Sub(*task.StartedAt).Seconds()
	task.ProcessingTime = &processingTime

	s.storage.Create(task)
	log.Printf("Task %s completed in %v", id, duration)
}

func generateID() string {
	return time.Now().Format("20060102150405") + "-" + randString(6)
}

func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[time.Now().UnixNano()%int64(len(letters))]
	}
	return string(b)
}

func (s *TaskService) GetAllTasks() []models.Task {
	return s.storage.GetAll()
}
