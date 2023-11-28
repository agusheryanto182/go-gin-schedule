package controllers

import "github.com/gin-gonic/gin"

type ScheduleController interface {
	GetData(c *gin.Context)
	AddSchedule(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}
