package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CheckIn(c *gin.Context)
}
