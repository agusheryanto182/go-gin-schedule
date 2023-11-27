package controllers

import (
	"net/http"
	"strconv"

	"github.com/agusheryanto182/go-schedule/helpers"
	"github.com/agusheryanto182/go-schedule/models/domain"
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

	findByEmail, _ := self.UserService.FindByEmail(email)

	_ = c.ShouldBindJSON(&input)
	if input.Title == "" {
		response := helpers.APIResponseFailed("Bad Request", "Title is empty")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Day == "" {
		response := helpers.APIResponseFailed("Bad Request", "Day is empty")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.UserId = findByEmail.UserId

	result, err := self.ScheduleService.Create(input)
	if err != nil {
		response := helpers.APIResponse("Bad Request", "Error on create", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success", "Success", result)

	c.JSON(http.StatusOK, response)
}

func (self *ScheduleControllerImpl) Edit(c *gin.Context) {
	var input web.ScheduleUpdateRequest

	_ = c.ShouldBindJSON(&input)

	email := c.Query("email")
	id, _ := strconv.Atoi(c.Query("id"))

	checkEmail, _ := self.UserService.FindByEmail(email)

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

	updated, _ := self.ScheduleService.Update(input)

	response := helpers.APIResponse("Success", "Success", updated)

	c.JSON(http.StatusOK, response)

}

func (self *ScheduleControllerImpl) Delete(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(domain.User)

	email := c.Query("email")
	id, _ := strconv.Atoi(c.Query("id"))

	checkSchedule, _ := self.ScheduleService.FindById(id)

	if checkSchedule.ScheduleId == 0 {
		response := helpers.APIResponse("No schedule on that ID", "Bad Request", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if currentUser.Email != email {
		response := helpers.APIResponse("Forbiden", "Denied", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := self.ScheduleService.Delete(id)
	if err != nil {
		response := helpers.APIResponse("Delete Schedule is failed", "Not Found", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success", "Success", nil)
	c.JSON(http.StatusOK, response)

}

func (self *ScheduleControllerImpl) GetAll(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(domain.User)

	email := c.Query("email")

	if currentUser.Email != email {
		response := helpers.APIResponse("Forbiden", "Denied", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	listSchedule := self.ScheduleService.FindAll()
	response := helpers.APIResponse("Success", "Success", listSchedule)
	c.JSON(http.StatusOK, response)
}

func (self *ScheduleControllerImpl) GetById(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(domain.User)

	email := c.Query("email")

	day := c.Query("day")

	if currentUser.Email != email {
		response := helpers.APIResponse("Forbiden", "Denied", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	checkDay, _ := self.ScheduleService.FindByDay(day)
	if checkDay.ScheduleId == 0 {
		response := helpers.APIResponse("No found schedule on that day", "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("Success", "Success", checkDay)
	c.JSON(http.StatusOK, response)
}

func NewScheduleController(ScheduleService services.ScheduleService, UserService services.UserService) ScheduleController {
	return &ScheduleControllerImpl{ScheduleService: ScheduleService, UserService: UserService}
}
