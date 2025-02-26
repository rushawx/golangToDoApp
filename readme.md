# Go ToDo Application

This is a simple ToDo application written in Go. It provides a RESTful API to manage tasks, including creating, updating, retrieving, and deleting tasks.

## Features

- Create a new task
- Get all tasks
- Get a task by ID
- Update a task by ID
- Delete a task by ID

## Getting Started

### Prerequisites

- Go 1.16 or later
- Git

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/golangToDo.git
    cd golangToDo
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Install `swag` for generating Swagger documentation:

    ```sh
    go get -u github.com/swaggo/swag/cmd/swag
    ```

### Running the Application

1. Generate Swagger documentation:

    ```sh
    swag init
    ```

2. Run the application:

    ```sh
    go run cmd/main.go
    ```

3. The application will start on `http://localhost:8080`.

### API Documentation

The API documentation is available via Swagger UI. After running the application, navigate to `http://localhost:8080/swagger/index.html` to view and interact with the API documentation.

### API Endpoints

- `GET /tasks`: Get all tasks
- `GET /tasks/{id}`: Get a task by ID
- `POST /tasks`: Create a new task
- `PUT /tasks/{id}`: Update a task by ID
- `DELETE /tasks/{id}`: Delete a task by ID

### Example Task JSON

```json
{
  "ID": "string",
  "Title": "string",
  "Description": "string",
  "Done": false,
  "ToDo": "2025-02-26T00:00:00Z",
  "CreatedAt": "2025-02-26T00:00:00Z",
  "UpdatedAt": "2025-02-26T00:00:00Z",
  "DeletedAt": "2025-02-26T00:00:00Z"
}
```

### Project Structure

```shell
golangToDo/
├── cmd/
│   └── main.go
├── internal/
│   └── todo/
│       ├── handler.go
│       ├── model.go
│       └── repository.go
├── pkg/
│   └── db/
│       └── db.go
├── docs/
│   └── swagger.yaml
├── go.mod
└── go.sum
```