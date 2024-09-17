package handler

import (
	"encoding/json"
	"go-psql/taskmanager/service"
	"net/http"
)

type TaskHandler struct {
	TaskService service.TaskService
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if task.Title == "" || task.Description == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if err := h.TaskService.CreateTask(task.Title, task.Description); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (h *TaskHandler) ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.TaskService.ListTasks()

	if err != nil {
		http.Error(w, "Failed to list tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
