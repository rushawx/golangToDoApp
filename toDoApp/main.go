package main

import (
	"fmt"
	"net/http"
	"toDo/internal/todo"
	"toDo/pkg/db"

	_ "toDo/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			ToDo API
//	@version		1.0
//	@description	This is a sample ToDo server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/
func main() {
	db := db.NewDb()

	router := http.NewServeMux()

	taskRepository := todo.NewTaskRepository(db)

	todo.NewTaskHandler(router, taskRepository)

	router.Handle("/swagger/", httpSwagger.WrapHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
