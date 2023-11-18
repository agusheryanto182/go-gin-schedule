package controllers

import (
	"net/http"
	"strconv"

	"github.com/agusheryanto182/go-todo/helpers"
	"github.com/agusheryanto182/go-todo/models/web"
	"github.com/agusheryanto182/go-todo/services"
	"github.com/gin-gonic/gin"
)

type TodoControllerImpl struct {
	todoService services.TodoService
}

func (h *TodoControllerImpl) Create(c *gin.Context) {
	var input web.TodoCreateRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, "title or activity_group_id cannot be null", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.Title == "" {
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, "title cannot be null", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.ActivityGroupId == 0 {
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, "activity_group_id cannot be null", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newTodo, err := h.todoService.Create(input)
	if err != nil {
		response := helpers.APIResponse("Bad Request", http.StatusBadRequest, "title or activity_group_id cannot be null", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success", http.StatusOK, "Success", helpers.TodoFormatter(newTodo))

	c.JSON(http.StatusOK, response)
}

func (h *TodoControllerImpl) Update(c *gin.Context) {
	var input web.TodoUpdateRequest
	todoId, _ := strconv.Atoi(c.Param("todoId"))

	checkTodo, err := h.todoService.GetById(todoId)
	if err != nil {
		response := helpers.APIResponse("Todo with ID "+strconv.Itoa(todoId)+" Not Found", http.StatusNotFound, " Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	if checkTodo.TodoId <= 0 {
		response := helpers.APIResponse("Todo with ID "+strconv.Itoa(todoId)+" Not Found", http.StatusNotFound, " Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helpers.APIResponse("Update todo is failed 1", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.TodoId = todoId

	updatedTodo, err := h.todoService.Update(input)
	if err != nil {
		response := helpers.APIResponse("Update todo is failed 2", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helpers.APIResponse("Success", http.StatusOK, "Success", helpers.TodoFormatter(updatedTodo))
	c.JSON(http.StatusOK, response)

}

func (h *TodoControllerImpl) Delete(c *gin.Context) {
	todoId, _ := strconv.Atoi(c.Param("todoId"))

	checkTodo, err := h.todoService.GetById(todoId)
	if err != nil {
		response := helpers.APIResponse("Todo with ID "+strconv.Itoa(todoId)+" Not Found", http.StatusNotFound, " Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	if checkTodo.TodoId <= 0 {
		response := helpers.APIResponse("Todo with ID "+strconv.Itoa(todoId)+" Not Found", http.StatusNotFound, " Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = h.todoService.Delete(todoId)
	if err != nil {
		response := helpers.APIResponse("Delete todo is failed", http.StatusBadRequest, " Not Found", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("Success", http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)

}

func (h *TodoControllerImpl) FindById(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("todoId"))
	if todoId <= 0 {
		response := helpers.APIResponse("Todo with ID "+strconv.Itoa(todoId)+" Not Found", http.StatusNotFound, " Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	todo, err := h.todoService.GetById(todoId)
	if err != nil {
		response := helpers.APIResponse("Todo with ID "+strconv.Itoa(todoId)+" Not Found", http.StatusNotFound, " Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	if todo.TodoId <= 0 {
		response := helpers.APIResponse("Todo with ID "+strconv.Itoa(todoId)+" Not Found", http.StatusNotFound, " Not Found", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helpers.APIResponse("Success", http.StatusOK, "Success", helpers.TodoFormatter(todo))
	c.JSON(http.StatusOK, response)

}

func (h *TodoControllerImpl) FindAll(c *gin.Context) {
	activityGroupId, _ := strconv.Atoi(c.Query("activity_group_id"))
	if activityGroupId == 0 {
		todo, _ := h.todoService.GetAll()
		response := helpers.APIResponse("Success", http.StatusOK, "Success", helpers.ToTodoResponses(todo))
		c.JSON(http.StatusOK, response)
	} else {
		todos, err := h.todoService.GetAllByActivityGroupId(activityGroupId)
		if err != nil {
			response := helpers.APIResponse("Something error", http.StatusBadRequest, "Bad Request", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helpers.APIResponse("Success", http.StatusOK, "Success", helpers.ToTodoResponses(todos))
		c.JSON(http.StatusOK, response)
	}
}

func NewTodoController(todoService services.TodoService) TodoController {
	return &TodoControllerImpl{todoService: todoService}
}
