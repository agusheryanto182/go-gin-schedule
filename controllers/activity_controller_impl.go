package controllers

import (
	"net/http"
	"strconv"

	"github.com/agusheryanto182/go-todo/helpers"
	"github.com/agusheryanto182/go-todo/models/web"
	"github.com/agusheryanto182/go-todo/services"
	"github.com/gin-gonic/gin"
)

type ActivityControllerImpl struct {
	activityService services.ActivityService
}

func (h *ActivityControllerImpl) Create(c *gin.Context) {
	var input web.ActivityCreateRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, "title cannot be null", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newActivity, err := h.activityService.Create(input)
	if err != nil {
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, "title cannot be null", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success", http.StatusOK, "Success", helpers.ToActivityResponse(newActivity))

	c.JSON(http.StatusOK, response)
}

func (h *ActivityControllerImpl) Update(c *gin.Context) {
	var input web.ActivityUpdateRequest
	activityId, _ := strconv.Atoi(c.Param("activityId"))

	checkActivity, err := h.activityService.GetById(activityId)
	if err != nil {
		response := helpers.APIResponse("Activity with ID "+strconv.Itoa(activityId)+" Not Found", http.StatusNotFound, "Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	if checkActivity.ActivityId <= 0 {
		response := helpers.APIResponse("Activity with ID "+strconv.Itoa(activityId)+" Not Found", http.StatusNotFound, "Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helpers.APIResponse("Bad request", http.StatusBadRequest, "Title cannot be null", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.ActivityId = activityId

	updatedActivity, err := h.activityService.Update(input)
	if err != nil {
		response := helpers.APIResponse("Bad request", http.StatusBadRequest, "Title cannot be null", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success", http.StatusOK, "Success", helpers.ToActivityResponse(updatedActivity))
	c.JSON(http.StatusOK, response)
}

func (h *ActivityControllerImpl) Delete(c *gin.Context) {
	activityId, _ := strconv.Atoi(c.Param("activityId"))

	checkActivity, err := h.activityService.GetById(activityId)
	if err != nil {
		response := helpers.APIResponse("Activity with ID "+strconv.Itoa(activityId)+" Not Found", http.StatusNotFound, "Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	if checkActivity.ActivityId <= 0 {
		response := helpers.APIResponse("Activity with ID "+strconv.Itoa(activityId)+" Not Found", http.StatusNotFound, "Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = h.activityService.Delete(activityId)
	if err != nil {
		response := helpers.APIResponse("Delete Activity is failed", http.StatusBadRequest, "Not Found", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success", http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)

}

func (h *ActivityControllerImpl) FindById(c *gin.Context) {
	activityId, _ := strconv.Atoi(c.Param("activityId"))

	if activityId <= 0 {
		response := helpers.APIResponse("Activity with ID "+strconv.Itoa(activityId)+" Not Found", http.StatusNotFound, "Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	activity, err := h.activityService.GetById(activityId)
	if err != nil {
		response := helpers.APIResponse("Activity with ID "+strconv.Itoa(activityId)+" Not Found", http.StatusNotFound, "Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	if activity.ActivityId <= 0 {
		response := helpers.APIResponse("Activity with ID "+strconv.Itoa(activityId)+" Not Found", http.StatusNotFound, "Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helpers.APIResponse("Success", http.StatusOK, "Success", helpers.ToActivityResponse(activity))
	c.JSON(http.StatusOK, response)

}

func (h *ActivityControllerImpl) FindAll(c *gin.Context) {
	activity, _ := h.activityService.GetAll()
	response := helpers.APIResponse("Success", http.StatusOK, "Success", helpers.ToActivityResponses(activity))
	c.JSON(http.StatusOK, response)
}

func NewActivityController(activityService services.ActivityService) ActivityController {
	return &ActivityControllerImpl{activityService: activityService}
}
