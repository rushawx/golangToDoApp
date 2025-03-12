package todo

import (
	"toDo/pkg/db"

	"gorm.io/gorm/clause"
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
	result := tr.Database.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (tr *TaskRepository) GetTask(id string) (*Task, error) {
	data := Task{}
	idBytes := []byte(id)
	result := tr.Database.First(&data, "task_id = ?", idBytes)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (tr *TaskRepository) GetTasks() ([]*Task, error) {
	var tasks []*Task
	result := tr.Database.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (tr *TaskRepository) UpdateTask(task *Task) (*Task, error) {
	result := tr.Database.Clauses(clause.Returning{}).Where("task_id = ?", task.TaskID).Updates(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (tr *TaskRepository) DeleteTask(id string) error {
	task, err := tr.GetTask(id)
	if err != nil {
		return err
	}
	result := tr.Database.Delete(task)
	return result.Error
}
