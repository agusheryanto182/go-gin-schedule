package controllers

import (
	"fmt"
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

	if !helpers.IsValidEmail(email) {
		response := helpers.APIResponseFailed("Bad Request", "Invalid email")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	findByEmail, _ := self.UserService.FindByEmail(email)
	if findByEmail.UserId == 0 {
		response := helpers.APIResponseFailed("Not Found", "Email is not found")
		c.JSON(http.StatusNotFound, response)
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

	if !helpers.IsDayValid(input.Day) {
		response := helpers.APIResponseFailed("Bad Request", "Day is invalid")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.UserId = findByEmail.UserId

	result, _ := self.ScheduleService.Create(input)

	response := helpers.APIResponse("Success", "Success", result)

	c.JSON(http.StatusCreated, response)
}

func (self *ScheduleControllerImpl) Edit(c *gin.Context) {
	var input web.ScheduleUpdateRequest

	_ = c.ShouldBindJSON(&input)
	if input.Title == "" {
		response := helpers.APIResponseFailed("Bad Request", "Title is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	email := c.Query("email")
	if email == "" {
		response := helpers.APIResponseFailed("Bad Request", "Email is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !helpers.IsValidEmail(email) {
		response := helpers.APIResponseFailed("Bad Request", "Invalid email")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id, _ := strconv.Atoi(c.Query("id"))

	checkEmail, _ := self.UserService.FindByEmail(email)
	if checkEmail.UserId == 0 {
		response := helpers.APIResponseFailed("Not Found", "Email is not found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	checkSchedule, _ := self.ScheduleService.FindById(id)

	if checkSchedule.ScheduleId == 0 {
		response := helpers.APIResponseFailed("Not Found", "Schedule with ID "+strconv.Itoa(id)+" Not Found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	if checkSchedule.UserId != checkEmail.UserId {
		response := helpers.APIResponseFailed("Forbidden", "Access denied!")
		c.JSON(403, response)
		return
	}

	input.ScheduleId = checkSchedule.ScheduleId
	input.UserId = checkEmail.UserId

	updated, err := self.ScheduleService.Update(input)
	if err != nil {
		response := helpers.APIResponseFailed("Bad Request", "Error")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(updated.Title)

	response := helpers.APIResponse("Success", "Success", updated)

	c.JSON(http.StatusCreated, response)

}

func (self *ScheduleControllerImpl) Delete(c *gin.Context) {

	email := c.Query("email")
	if email == "" {
		response := helpers.APIResponseFailed("Bad Request", "Email is required")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !helpers.IsValidEmail(email) {
		response := helpers.APIResponseFailed("Bad Request", "Invalid email")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	checkEmail, _ := self.UserService.FindByEmail(email)
	if checkEmail.UserId == 0 {
		response := helpers.APIResponseFailed("Not Found", "Email is not found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	id := c.Query("id")
	intId, _ := strconv.Atoi(id)

	if id == "null" {
		_ = self.ScheduleService.Delete(checkEmail.UserId)

		response := helpers.APIResponse("Success", "Success", map[string]interface{}{})
		c.JSON(http.StatusOK, response)
		return
	}

	checkSchedule, _ := self.ScheduleService.FindById(intId)

	if checkSchedule.ScheduleId == 0 {
		response := helpers.APIResponseFailed("Not Found", "Schedule with ID "+strconv.Itoa(intId)+" Not Found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	if checkEmail.UserId != checkSchedule.UserId {
		response := helpers.APIResponseFailed("Forbidden", "Access denied!")
		c.JSON(http.StatusForbidden, response)
		return
	}

	_ = self.ScheduleService.Delete(intId)

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

	if !helpers.IsValidEmail(email) {
		response := helpers.APIResponseFailed("Bad Request", "Invalid email")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	checkEmail, _ := self.UserService.FindByEmail(email)
	if checkEmail.UserId == 0 {
		response := helpers.APIResponseFailed("Not Found", "Email is not found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	day := c.Query("day")
	if day == "" {
		monday, _ := self.ScheduleService.FindByDayAndUserId("monday", checkEmail.UserId)
		tuesday, _ := self.ScheduleService.FindByDayAndUserId("tuesday", checkEmail.UserId)
		wednesday, _ := self.ScheduleService.FindByDayAndUserId("wednesday", checkEmail.UserId)
		thursday, _ := self.ScheduleService.FindByDayAndUserId("thursday", checkEmail.UserId)
		friday, _ := self.ScheduleService.FindByDayAndUserId("friday", checkEmail.UserId)

		response := helpers.ResponseGetAll{
			Status:  "Success",
			Message: "Success",
			Data: map[string][]web.ScheduleResponse{
				"monday":    monday,
				"tuesday":   tuesday,
				"wednesday": wednesday,
				"thursday":  thursday,
				"friday":    friday,
			},
		}

		if len(monday) == 0 {
			response.Data["monday"] = []web.ScheduleResponse{}
		}

		if len(tuesday) == 0 {
			response.Data["tuesday"] = []web.ScheduleResponse{}
		}

		if len(wednesday) == 0 {
			response.Data["wednesday"] = []web.ScheduleResponse{}
		}

		if len(thursday) == 0 {
			response.Data["thursday"] = []web.ScheduleResponse{}
		}

		if len(friday) == 0 {
			response.Data["friday"] = []web.ScheduleResponse{}
		}
		c.JSON(http.StatusOK, response)
		return
	}

	if !helpers.IsDayValid(day) {
		response := helpers.APIResponseFailed("Bad Request", "Day is invalid")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	checkDay, _ := self.ScheduleService.FindByDay(day)

	response := helpers.APIResponse("Success", "Success", checkDay)
	c.JSON(http.StatusOK, response)

}

func NewScheduleController(ScheduleService services.ScheduleService, UserService services.UserService) ScheduleController {
	return &ScheduleControllerImpl{ScheduleService: ScheduleService, UserService: UserService}
}
