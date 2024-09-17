package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type TaskStore interface {
	SaveTask(title, description string) error
	GetAllTasks() ([]Task, error)
}

type SQLTaskStore struct {
	DB *sql.DB
}

func (s *SQLTaskStore) SaveTask(title, description string) error {
	_, err := s.DB.Exec("INSERT INTO tasks (title, description) VALUES ($1, $2)", title, description)
	return err
}

func (s *SQLTaskStore) GetAllTasks() ([]Task, error) {
	rows, err := s.DB.Query("SELECT id, title, description, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
