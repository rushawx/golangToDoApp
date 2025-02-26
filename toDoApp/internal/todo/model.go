package todo

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Done        bool
	ToDo        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(title string, description string, todo time.Time) *Task {
	return &Task{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Done:        false,
		ToDo:        todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
