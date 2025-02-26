package todo

import (
	"errors"
	"toDo/pkg/db"
)

type TaskRepository struct {
	Database *db.Db
}

func NewTaskRepository(db *db.Db) *TaskRepository {
	return &TaskRepository{
		Database: db,
	}
}

func (tr *TaskRepository) CreateTask(task *Task) (*Task, error) {
	tr.Database.Tasks[task.ID] = db.TaskDb{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Done:        task.Done,
		ToDo:        task.ToDo,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	return task, nil
}

func (tr *TaskRepository) GetTask(id string) (*Task, error) {
	task, ok := tr.Database.Tasks[id]
	if !ok {
		return nil, errors.New("Task not found")
	}
	data := &Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Done:        task.Done,
		ToDo:        task.ToDo,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	return data, nil
}

func (tr *TaskRepository) GetTasks() ([]*Task, error) {
	var tasks []*Task
	for _, task := range tr.Database.Tasks {
		tasks = append(tasks, &Task{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
			ToDo:        task.ToDo,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}
	return tasks, nil
}

func (tr *TaskRepository) UpdateTask(task *Task) (*Task, error) {
	_, ok := tr.Database.Tasks[task.ID]
	if !ok {
		return nil, errors.New("Task not found")
	}
	tr.Database.Tasks[task.ID] = db.TaskDb{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Done:        task.Done,
		ToDo:        task.ToDo,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	return task, nil
}

func (tr *TaskRepository) DeleteTask(id string) error {
	_, ok := tr.Database.Tasks[id]
	if !ok {
		return errors.New("Task not found")
	}
	delete(tr.Database.Tasks, id)
	return nil
}
