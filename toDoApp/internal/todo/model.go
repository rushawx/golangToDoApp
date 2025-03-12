package todo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskID      []byte    `json:"task_id" gorm:"uniqueIndex"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	ToDo        time.Time `json:"todo"`
}

func NewTask(title string, description string, todo time.Time) *Task {
	u := uuid.New()
	uBytes := u[:]

	return &Task{
		TaskID:      uBytes,
		Title:       title,
		Description: description,
		Done:        false,
		ToDo:        todo,
	}
}
