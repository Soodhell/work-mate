package storage

import (
	"sync"
	"task-manager/internal/models"
)

type TaskStorage interface {
	Create(task models.Task)
	Get(id string) (models.Task, bool)
	Delete(id string)
	GetAll() []models.Task
}

type MemoryStorage struct {
	tasks map[string]models.Task
	mu    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[string]models.Task),
	}
}

func (s *MemoryStorage) Create(task models.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tasks[task.ID] = task
}

func (s *MemoryStorage) Get(id string) (models.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	task, exists := s.tasks[id]
	return task, exists
}

func (s *MemoryStorage) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tasks, id)
}

func (s *MemoryStorage) GetAll() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()
	tasks := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}
