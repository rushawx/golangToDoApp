package todo

import (
	"time"
)

type TaskRequest struct {
	Title       string
	Description string
	ToDo        time.Time
}
