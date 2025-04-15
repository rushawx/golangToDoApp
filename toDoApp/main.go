package main

import (
	"fmt"
	"net/http"
	"toDo/configs"
	"toDo/internal/auth"
	"toDo/internal/todo"
	"toDo/internal/user"
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
func App() http.Handler {
	conf := configs.DefaultConfig()
	db := db.NewDb(conf)

	router := http.NewServeMux()

	taskRepository := todo.NewTaskRepository(db)
	userRepository := user.NewUserRepository(db)
	sessionRepository := user.NewSessionRepository(db)

	authService := auth.NewAuthService(userRepository, sessionRepository)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf, AuthService: authService})
	todo.NewTaskHandler(router, todo.TaskHandlerDeps{TaskRepository: taskRepository, Config: conf})

	router.Handle("/swagger/", httpSwagger.WrapHandler)

	return router
}

func main() {
	router := App()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
