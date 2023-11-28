package app

import (
	"github.com/agusheryanto182/go-schedule/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController controllers.UserController, scheduleController controllers.ScheduleController) *gin.Engine {
	router := gin.Default()

	router.POST("/checkin", userController.CheckIn)

	schedule := router.Group("/schedule")
	schedule.GET("", scheduleController.GetData)
	schedule.POST("", scheduleController.AddSchedule)
	schedule.PATCH("", scheduleController.Edit)
	schedule.DELETE("", scheduleController.Delete)

	return router
}
