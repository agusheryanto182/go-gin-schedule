package main

import (
	"fmt"
	"net/http"

	"github.com/agusheryanto182/go-schedule/app"
	"github.com/agusheryanto182/go-schedule/controllers"
	"github.com/agusheryanto182/go-schedule/middlewares"
	"github.com/agusheryanto182/go-schedule/repositories"
	"github.com/agusheryanto182/go-schedule/services"
	"github.com/agusheryanto182/go-schedule/utils/log"
	"github.com/go-playground/validator/v10"
)

func InitServer() *http.Server {
	conf := app.NewConfig()
	db := app.NewDB(conf)
	validate := validator.New()

	UserRepository := repositories.NewUserRepository(db)
	UserService := services.NewUserService(UserRepository, validate)
	userController := controllers.NewUserController(UserService)

	scheduleRepository := repositories.NewScheduleRepository(db)
	scheduleService := services.NewScheduleService(scheduleRepository, validate)
	scheduleController := controllers.NewScheduleController(scheduleService, UserService)

	r := app.NewRouter(userController, scheduleController)
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
