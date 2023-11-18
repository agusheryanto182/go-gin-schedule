package services

import (
	"github.com/agusheryanto182/go-todo/models/domain"
	"github.com/agusheryanto182/go-todo/models/web"
)

type TodoService interface {
	Create(input web.TodoCreateRequest) (web.TodoResponse, error)
	Update(input web.TodoUpdateRequest) (web.TodoResponse, error)
	Delete(Id int) error
	GetById(Id int) (web.TodoResponse, error)
	GetAll() ([]domain.Todo, error)
	GetAllByActivityGroupId(activityGroupId int) ([]domain.Todo, error)
}
