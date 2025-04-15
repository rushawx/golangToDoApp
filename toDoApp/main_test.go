package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"toDo/internal/todo"

	"github.com/google/uuid"
)

func TestCreateTask(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	task := todo.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		ToDo:        time.Now(),
		Done:        false,
	}

	data, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/tasks", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Failed to create POST request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ts.Client().Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		t.Fatalf("Expected status code 201, got %d", resp.StatusCode)
	}

	var createdTask todo.Task
	err = json.NewDecoder(resp.Body).Decode(&createdTask)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
	if createdTask.TaskID == uuid.Nil {
		t.Fatalf("Expected task ID to be non-empty, got %s", createdTask.TaskID)
	}
	if createdTask.Title != task.Title {
		t.Fatalf("Expected task title %s, got %s", task.Title, createdTask.Title)
	}
	if createdTask.Description != task.Description {
		t.Fatalf("Expected task description %s, got %s", task.Description, createdTask.Description)
	}
	if createdTask.Done != task.Done {
		t.Fatalf("Expected task status %v, got %v", task.Done, createdTask.Done)
	}
}
func TestGetTask(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	task := todo.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		ToDo:        time.Now(),
		Done:        false,
	}

	data, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/tasks", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Failed to create POST request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ts.Client().Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		t.Fatalf("Expected status code 201, got %d", resp.StatusCode)
	}

	var createdTask todo.Task
	err = json.NewDecoder(resp.Body).Decode(&createdTask)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	resp, err = ts.Client().Get(ts.URL + "/tasks/" + createdTask.TaskID.String())
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
	if task.TaskID == uuid.Nil {
		t.Fatalf("Expected task ID to be non-empty, got %s", task.TaskID)
	}
	if task.Title == "" {
		t.Fatalf("Expected task title to be non-empty, got %s", task.Title)
	}
	if task.Description == "" {
		t.Fatalf("Expected task description to be non-empty, got %s", task.Description)
	}
	if !task.Done && task.Done {
		t.Fatalf("Expected task status to be either true or false, got %v", task.Done)
	}
}
func TestUpdateTask(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	task := todo.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		ToDo:        time.Now(),
		Done:        false,
	}

	data, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/tasks", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Failed to create POST request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ts.Client().Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		t.Fatalf("Expected status code 201, got %d", resp.StatusCode)
	}

	var createdTask todo.Task
	err = json.NewDecoder(resp.Body).Decode(&createdTask)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	updatedData, err := json.Marshal(createdTask)
	if err != nil {
		t.Fatalf("Failed to marshal updated task: %v", err)
	}
	req, err = http.NewRequest(http.MethodPut, ts.URL+"/tasks/"+createdTask.TaskID.String(), bytes.NewReader(updatedData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatalf("Failed to create PUT request: %v", err)
	}
	resp, err = ts.Client().Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		t.Fatalf("Expected status code 201, got %d", resp.StatusCode)
	}

	var updatedTask todo.Task
	err = json.NewDecoder(resp.Body).Decode(&updatedTask)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
	if updatedTask.TaskID == uuid.Nil {
		t.Fatalf("Expected task ID to be non-empty, got %s", updatedTask.TaskID)
	}
	if updatedTask.Title != task.Title {
		t.Fatalf("Expected task title %s, got %s", task.Title, updatedTask.Title)
	}
	if updatedTask.Description != task.Description {
		t.Fatalf("Expected task description %s, got %s", task.Description, updatedTask.Description)
	}
	if updatedTask.Done != task.Done {
		t.Fatalf("Expected task status %v, got %v", task.Done, updatedTask.Done)
	}
}

func TestGetTasks(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/tasks")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}

	var tasks []todo.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}
	if len(tasks) == 0 {
		t.Fatalf("Expected at least one task, got %d", len(tasks))
	}

	for _, task := range tasks {
		if task.TaskID == uuid.Nil {
			t.Fatalf("Expected task ID to be non-empty, got %s", task.TaskID)
		}
		if task.Title == "" {
			t.Fatalf("Expected task title to be non-empty, got %s", task.Title)
		}
		if task.Description == "" {
			t.Fatalf("Expected task description to be non-empty, got %s", task.Description)
		}
		if !task.Done && task.Done {
			t.Fatalf("Expected task status to be either true or false, got %v", task.Done)
		}
	}
}

func TestDeleteTask(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	task := todo.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		ToDo:        time.Now(),
		Done:        false,
	}

	data, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Failed to marshal task: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, ts.URL+"/tasks", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Failed to create POST request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ts.Client().Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		t.Fatalf("Expected status code 201, got %d", resp.StatusCode)
	}

	var createdTask todo.Task
	err = json.NewDecoder(resp.Body).Decode(&createdTask)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	req, err = http.NewRequest("DELETE", ts.URL+"/tasks/"+createdTask.TaskID.String(), nil)
	if err != nil {
		t.Fatalf("Failed to create DELETE request: %v", err)
	}
	resp, err = ts.Client().Do(req)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		t.Fatalf("Expected status code 204, got %d", resp.StatusCode)
	}
}
