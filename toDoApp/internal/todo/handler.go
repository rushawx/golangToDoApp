package todo

import (
	"encoding/json"
	"net/http"
	"time"
)

type TaskHandler struct {
	TaskRepository *TaskRepository
}

func NewTaskHandler(router *http.ServeMux, tr *TaskRepository) {
	th := &TaskHandler{
		TaskRepository: tr,
	}
	router.HandleFunc("GET /tasks", th.GetTasks())
	router.HandleFunc("GET /tasks/{id}", th.GetTask())
	router.HandleFunc("POST /tasks", th.CreateTask())
	router.HandleFunc("PUT /tasks/{id}", th.UpdateTask())
	router.HandleFunc("DELETE /tasks/{id}", th.DeleteTask())
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

// @Summary		Get a task by ID
// @Description	Get a task by ID
// @Tags			tasks
// @Produce		json
// @Param			id	path		string	true	"Task ID"
// @Success		200	{object}	Task
// @Router			/tasks/{id} [get]
func (th *TaskHandler) GetTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		data, err := th.TaskRepository.GetTask(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

// @Summary		Create a new task
// @Description	Create a new task with the input payload
// @Tags			tasks
// @Accept			json
// @Produce		json
// @Param			task	body		TaskRequest	true	"Task to create"
// @Success		200		{object}	Task
// @Router			/tasks [post]
func (th *TaskHandler) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input TaskRequest
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := NewTask(input.Title, input.Description, input.ToDo)
		task, err := th.TaskRepository.CreateTask(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)

	}
}

// @Summary		Update a task by ID
// @Description	Update a task by ID
// @Tags			tasks
// @Accept			json
// @Produce		json
// @Param			id		path		string	true	"Task ID"
// @Param			task	body		Task	true	"Task to update"
// @Success		200		{object}	Task
// @Router			/tasks/{id} [put]
func (th *TaskHandler) UpdateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var task Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		taskDb, err := th.TaskRepository.GetTask(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		task.ID = id
		task.CreatedAt = taskDb.CreatedAt
		task.UpdatedAt = time.Now()
		if task.Title == "" {
			task.Title = taskDb.Title
		}
		if task.Description == "" {
			task.Description = taskDb.Description
		}
		if task.ToDo.IsZero() {
			task.ToDo = taskDb.ToDo
		}
		data, err := th.TaskRepository.UpdateTask(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

// @Summary		Delete a task by ID
// @Description	Delete a task by ID
// @Tags			tasks
// @Param			id	path	string	true	"Task ID"
// @Success		204
// @Router			/tasks/{id} [delete]
func (th *TaskHandler) DeleteTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		err := th.TaskRepository.DeleteTask(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
