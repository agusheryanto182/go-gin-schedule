package helpers

import (
	"time"

	"github.com/agusheryanto182/go-todo/models/domain"
	"github.com/agusheryanto182/go-todo/models/web"
)

func ToActivityResponse(activity domain.Activity) web.ActivityResponse {
	var deletedAt *time.Time
	if !activity.DeletedAt.Time.IsZero() {
		deletedAt = &activity.DeletedAt.Time
	}

	return web.ActivityResponse{
		ActivityId: activity.ActivityId,
		Email:      activity.Email,
		Title:      activity.Title,
		CreatedAt:  activity.CreatedAt,
		UpdatedAt:  activity.UpdatedAt,
		DeletedAt:  deletedAt,
	}
}

func ToActivityResponses(listActivity []domain.Activity) []web.ActivityResponse {
	var activityResponses []web.ActivityResponse
	for _, activity := range listActivity {
		activityResponses = append(activityResponses, ToActivityResponse(activity))
	}

	return activityResponses
}

func ToTodoResponse(todo domain.Todo) web.TodoResponse {
	var deletedAt *time.Time
	if !todo.DeletedAt.Time.IsZero() {
		deletedAt = &todo.DeletedAt.Time
	}

	return web.TodoResponse{
		TodoId:          todo.TodoId,
		ActivityGroupId: todo.ActivityGroupId,
		Title:           todo.Title,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		CreatedAt:       todo.CreatedAt,
		UpdatedAt:       todo.UpdatedAt,
		DeletedAt:       deletedAt,
	}
}

func ToTodoResponses(listTodo []domain.Todo) []web.TodoResponse {
	var todoResponses []web.TodoResponse
	for _, todo := range listTodo {
		todoResponses = append(todoResponses, ToTodoResponse(todo))
	}

	return todoResponses
}

func TodoFormatter(todo web.TodoResponse) web.TodoResponse {
	todoFormatter := web.TodoResponse{}
	todoFormatter.TodoId = todo.TodoId
	todoFormatter.Title = todo.Title
	todoFormatter.ActivityGroupId = todo.ActivityGroupId
	todoFormatter.IsActive = todo.IsActive
	todoFormatter.Priority = todo.Priority
	todoFormatter.UpdatedAt = todo.UpdatedAt
	todoFormatter.CreatedAt = todo.CreatedAt
	todoFormatter.DeletedAt = todo.DeletedAt

	return todoFormatter
}
