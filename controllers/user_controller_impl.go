package controllers

import (
	"net/http"

	"github.com/agusheryanto182/go-schedule/helpers"
	"github.com/agusheryanto182/go-schedule/models/web"
	"github.com/agusheryanto182/go-schedule/services"
	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func (self *UserControllerImpl) CheckIn(c *gin.Context) {
	var input web.UserCreateRequest

	err := c.ShouldBindJSON(&input)
	if input.Email == "" {
		response := helpers.APIResponseFailed("Bad Request", "Email is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !helpers.IsValidEmail(input.Email) {
		response := helpers.APIResponseFailed("Bad Request", "Invalid email")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	checkEmail, _ := self.UserService.FindByEmail(input.Email)
	if checkEmail.UserId != 0 {
		response := helpers.APIResponse("Success", "Success", checkEmail)

		c.JSON(http.StatusOK, response)
		return
	}
	if err != nil {
		response := helpers.APIResponseFailed("Bad Request", "Invalid email")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := self.UserService.Create(input)
	if err != nil {
		response := helpers.APIResponseFailed("Bad Request", "Bad request")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success", "Success", result)

	c.JSON(http.StatusOK, response)

}

func NewUserController(UserService services.UserService) UserController {
	return &UserControllerImpl{UserService: UserService}
}
