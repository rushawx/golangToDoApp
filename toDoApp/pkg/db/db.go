package db

import "time"

type TaskDb struct {
	ID          string
	Title       string
	Description string
	Done        bool
	ToDo        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Db struct {
	Tasks map[string]TaskDb
}

func NewDb() *Db {
	return &Db{
		Tasks: make(map[string]TaskDb),
	}
}
