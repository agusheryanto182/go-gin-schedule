package repositories

import (
	"github.com/agusheryanto182/go-todo/models/domain"
)

type TodoRepository interface {
	Save(todo domain.Todo) (domain.Todo, error)
	Update(todo domain.Todo) (domain.Todo, error)
	Delete(Id int) error
	FindById(Id int) (domain.Todo, error)
	FindAll() ([]domain.Todo, error)
	FindAllByActivityGroupId(activityGroupId int) ([]domain.Todo, error)
}
