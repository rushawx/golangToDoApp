package todo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskID      uuid.UUID `json:"task_id" gorm:"uniqueIndex"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	ToDo        time.Time `json:"todo"`
}

func generateUUID() uuid.UUID {
	return uuid.MustParse(uuid.New().String())
}

func NewTask(title string, description string, todo time.Time) *Task {
	return &Task{
		TaskID:      generateUUID(),
		Title:       title,
		Description: description,
		Done:        false,
		ToDo:        todo,
	}
}
