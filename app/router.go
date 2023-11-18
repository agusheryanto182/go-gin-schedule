package app

import (
	"github.com/agusheryanto182/go-todo/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(activityController controllers.ActivityController, todoController controllers.TodoController) *gin.Engine {
	router := gin.Default()

	todoRoute := router.Group("/todo-items")
	{
		todoRoute.POST("/", todoController.Create)
		todoRoute.GET("/:todoId", todoController.FindById)
		todoRoute.PATCH("/:todoId", todoController.Update)
		todoRoute.DELETE("/:todoId", todoController.Delete)
		todoRoute.GET("", todoController.FindAll)

	}

	activityRoute := router.Group("/activity-groups")
	{
		activityRoute.GET("/", activityController.FindAll)
		activityRoute.POST("/", activityController.Create)
		activityRoute.GET("/:activityId", activityController.FindById)
		activityRoute.PATCH("/:activityId", activityController.Update)
		activityRoute.DELETE("/:activityId", activityController.Delete)

	}

	return router
}
