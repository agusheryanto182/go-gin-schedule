package controllers

import (
	"net/http"
	"strconv"

	"github.com/agusheryanto182/go-schedule/helpers"
	"github.com/agusheryanto182/go-schedule/models/web"
	"github.com/agusheryanto182/go-schedule/services"
	"github.com/gin-gonic/gin"
)

type ScheduleControllerImpl struct {
	ScheduleService services.ScheduleService
	UserService     services.UserService
}

func (self *ScheduleControllerImpl) AddSchedule(c *gin.Context) {
	var input web.ScheduleCreateRequest

	email := c.Query("email")
	if email == "" {
		response := helpers.APIResponseFailed("Bad Request", "Email is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	findByEmail, _ := self.UserService.FindByEmail(email)
	if findByEmail.UserId == 0 {
		response := helpers.APIResponseFailed("Bad Request", "Invalid email")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_ = c.ShouldBindJSON(&input)
	if input.Title == "" {
		response := helpers.APIResponseFailed("Bad Request", "Title is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Day == "" {
		response := helpers.APIResponseFailed("Bad Request", "Day is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Day != "monday" && input.Day != "tuesday" && input.Day != "wednesday" && input.Title != "thursday" && input.Day != "friday" {
		response := helpers.APIResponseFailed("Bad Request", "Invalid day")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.UserId = findByEmail.UserId

	result, _ := self.ScheduleService.Create(input)

	response := helpers.APIResponse("Success", "Success", result)

	c.JSON(http.StatusOK, response)
}

func (self *ScheduleControllerImpl) Edit(c *gin.Context) {
	var input web.ScheduleUpdateRequest

	_ = c.ShouldBindJSON(&input)
	if input.Title == "" {
		response := helpers.APIResponseFailed("Bad Request", "Title is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Day == "" {
		response := helpers.APIResponseFailed("Bad Request", "Day is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	email := c.Query("email")
	if email == "" {
		response := helpers.APIResponseFailed("Bad Request", "Email is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))

	checkEmail, _ := self.UserService.FindByEmail(email)
	if checkEmail.UserId == 0 {
		response := helpers.APIResponseFailed("Bad Request", "Invalid email")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	checkSchedule, _ := self.ScheduleService.FindById(id)

	if checkSchedule.ScheduleId == 0 {
		response := helpers.APIResponseFailed("Not Found", "Schedule with ID "+strconv.Itoa(id)+" Not Found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	if checkSchedule.UserId != checkEmail.UserId {
		response := helpers.APIResponseFailed("Forbidden", "Access denied")
		c.JSON(http.StatusForbidden, response)
		return
	}

	if checkSchedule.ScheduleId != id {
		response := helpers.APIResponseFailed("Forbidden", "Access denied")
		c.JSON(http.StatusForbidden, response)
		return
	}

	input.ScheduleId = checkSchedule.ScheduleId
	input.UserId = checkEmail.UserId

	updated, _ := self.ScheduleService.Update(input)

	response := helpers.APIResponse("Success", "Success", updated)

	c.JSON(http.StatusOK, response)

}

func (self *ScheduleControllerImpl) Delete(c *gin.Context) {

	email := c.Query("email")
	if email == "" {
		response := helpers.APIResponseFailed("Bad Request", "Email is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	checkEmail, _ := self.UserService.FindByEmail(email)
	if checkEmail.UserId == 0 {
		response := helpers.APIResponseFailed("Not Found", "Email is not found")
		c.JSON(http.StatusNotFound, response)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	checkSchedule, _ := self.ScheduleService.FindById(id)

	if checkSchedule.ScheduleId == 0 {
		response := helpers.APIResponseFailed("Not Found", "Schedule with ID "+strconv.Itoa(id)+" Not Found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	if checkEmail.UserId != checkSchedule.UserId {
		response := helpers.APIResponseFailed("Forbidden", "Access Denied!")
		c.JSON(http.StatusForbidden, response)
		return
	}

	_ = self.ScheduleService.Delete(id)

	response := helpers.APIResponse("Success", "Success", map[string]interface{}{})
	c.JSON(http.StatusOK, response)

}

func (self *ScheduleControllerImpl) GetData(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		response := helpers.APIResponseFailed("Bad Request", "Email is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	checkEmail, _ := self.UserService.FindByEmail(email)
	if checkEmail.UserId == 0 {
		response := helpers.APIResponseFailed("Bad Request", "Invalid email")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	day := c.Query("day")
	if day == "" {
		listSchedule, _ := self.ScheduleService.FindByUserId(checkEmail.UserId)

		response := helpers.APIResponse("Success", "Success", listSchedule)
		c.JSON(http.StatusOK, response)
		return
	}

	checkDay, _ := self.ScheduleService.FindByDay(day)
	if day != "monday" && day != "tuesday" && day != "wednesday" && day != "thursday" && day != "friday" {
		response := helpers.APIResponseFailed("Bad Request", "Day is invalid")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("Success", "Success", checkDay)
	c.JSON(http.StatusOK, response)

}

func NewScheduleController(ScheduleService services.ScheduleService, UserService services.UserService) ScheduleController {
	return &ScheduleControllerImpl{ScheduleService: ScheduleService, UserService: UserService}
}
