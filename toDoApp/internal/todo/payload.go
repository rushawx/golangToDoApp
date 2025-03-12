package todo

import (
	"time"
)

type TaskCreateRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	ToDo        time.Time `json:"todo" validate:"required"`
}

type TaskUpdateRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ToDo        time.Time `json:"todo"`
	Done        bool      `json:"done"`
}
