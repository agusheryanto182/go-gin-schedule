package repositories

import "github.com/agusheryanto182/go-todo/models/domain"

type ActivityRepository interface {
	Save(activity domain.Activity) (domain.Activity, error)
	Update(activity domain.Activity) (domain.Activity, error)
	Delete(Id int) error
	FindById(Id int) (domain.Activity, error)
	FindAll() ([]domain.Activity, error)
}