package service

import "go-psql/taskmanager/store"

type TaskService struct {
	TaskStore store.TaskStore
}

func (s *TaskService) CreateTask(title, description string) error {
	return s.TaskStore.SaveTask(title, description)
}

func (s *TaskService) ListTasks() ([]store.Task, error) {
	return s.TaskStore.GetAllTasks()
}
