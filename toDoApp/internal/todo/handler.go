package todo

import (
	"net/http"
	"time"
	"toDo/configs"
	"toDo/pkg/middleware"
	"toDo/pkg/request"
	"toDo/pkg/response"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskHandlerDeps struct {
	TaskRepository *TaskRepository
	Config         *configs.Config
}

type TaskHandler struct {
	TaskRepository *TaskRepository
	Config         *configs.Config
}

func NewTaskHandler(router *http.ServeMux, deps TaskHandlerDeps) {
	th := &TaskHandler{
		TaskRepository: deps.TaskRepository,
	}
	router.HandleFunc("GET /tasks", th.GetTasks())
	router.HandleFunc("GET /tasks/{id}", th.GetTask())
	router.Handle("POST /tasks", middleware.IsAuthed(th.CreateTask(), deps.Config))
	router.Handle("PUT /tasks/{id}", middleware.IsAuthed(th.UpdateTask(), deps.Config))
	router.Handle("DELETE /tasks/{id}", middleware.IsAuthed(th.DeleteTask(), deps.Config))
}

// @Summary		Get all tasks
// @Description	Get all tasks
// @Tags			tasks
// @Produce		json
// @Success		200	{array}	Task
// @Router			/tasks [get]
func (th *TaskHandler) GetTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := th.TaskRepository.GetTasks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, data, http.StatusOK)
	}
}

// @Summary		Get a task by ID
// @Description	Get a task by ID
// @Tags			tasks
// @Produce		json
// @Param			id	path		string	true	"Task ID"
// @Success		200	{object}	Task
// @Failure		404	{string}	string	"Task not found"
// @Router			/tasks/{id} [get]
func (th *TaskHandler) GetTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		data, err := th.TaskRepository.GetTask(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, data, http.StatusOK)
	}
}

// @Summary		Create a new task
// @Description	Create a new task with the input payload
// @Tags			tasks
// @Accept			json
// @Produce		json
// @Param			task	body		TaskCreateRequest	true	"Task to create"
// @Success		201		{object}	Task
// @Failure		400		{string}	string	"Invalid input"
// @Failure		500		{string}	string	"Internal server error"
// @Router			/tasks [post]
func (th *TaskHandler) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[TaskCreateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := NewTask(body.Title, body.Description, body.ToDo)
		task, err := th.TaskRepository.CreateTask(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, task, http.StatusCreated)
	}
}

// @Summary		Update a task by ID
// @Description	Update a task by ID
// @Tags			tasks
// @Accept			json
// @Produce		json
// @Param			id		path		string	true	"Task ID"
// @Param			task	body		TaskUpdateRequest	true	"Task to update"
// @Success		200		{object}	Task
// @Failure		400		{string}	string	"Invalid input"
// @Failure		404		{string}	string	"Task not found"
// @Failure		500		{string}	string	"Internal server error"
// @Router			/tasks/{id} [put]
func (th *TaskHandler) UpdateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		body, err := request.HandleBody[TaskUpdateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task, err := th.TaskRepository.GetTask(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var title string
		if body.Title == "" {
			title = task.Title
		} else {
			title = body.Title
		}
		var description string
		if body.Description == "" {
			description = task.Description
		} else {
			description = body.Description
		}
		var todo time.Time
		if body.ToDo.IsZero() {
			todo = task.ToDo
		} else {
			todo = body.ToDo
		}
		id, err := uuid.Parse(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := th.TaskRepository.UpdateTask(&Task{
			Model:       gorm.Model{ID: uint(task.ID)},
			TaskID:      id,
			Title:       title,
			Description: description,
			ToDo:        todo,
			Done:        body.Done,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, data, http.StatusCreated)
	}
}

// @Summary		Delete a task by ID
// @Description	Delete a task by ID
// @Tags			tasks
// @Param			id	path	string	true	"Task ID"
// @Success		204
// @Failure		404	{string}	string	"Task not found"
// @Failure		500	{string}	string	"Internal server error"
// @Router			/tasks/{id} [delete]
func (th *TaskHandler) DeleteTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		err := th.TaskRepository.DeleteTask(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
