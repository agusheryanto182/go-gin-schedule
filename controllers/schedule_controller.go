package controllers

import "github.com/gin-gonic/gin"

type ScheduleController interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	AddSchedule(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
}
