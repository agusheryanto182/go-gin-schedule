package main

import (
	"fmt"
	"net/http"

	"github.com/agusheryanto182/go-todo/app"
	"github.com/agusheryanto182/go-todo/controllers"
	"github.com/agusheryanto182/go-todo/middlewares"
	"github.com/agusheryanto182/go-todo/repositories"
	"github.com/agusheryanto182/go-todo/services"
	"github.com/agusheryanto182/go-todo/utils/log"
	"github.com/go-playground/validator/v10"
)

func InitServer() *http.Server {
	conf := app.NewConfig()
	db := app.NewDB(conf)
	validate := validator.New()

	todoRepository := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepository, validate)
	todoController := controllers.NewTodoController(todoService)

	activityRepository := repositories.NewActivityRepository(db)
	activityService := services.NewActivityService(activityRepository, validate)
	activityController := controllers.NewActivityController(activityService)

	r := app.NewRouter(activityController, todoController)
	err := r.Run()
	if err != nil {
		fmt.Println("Error on the route run")
	}
	authMiddleware := middlewares.NewAuthMiddleware(r, conf)

	logger := log.NewLogger(conf)

	recorderMiddleware := middlewares.NewRecorderMiddleware(authMiddleware, logger)
	server := app.NewServer(recorderMiddleware, conf)
	return server

}

func main() {
	server := InitServer()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
