package services

import (
	"github.com/agusheryanto182/go-todo/models/domain"
	"github.com/agusheryanto182/go-todo/models/web"
)

type ActivityService interface {
	Create(input web.ActivityCreateRequest) (domain.Activity, error)
	Update(input web.ActivityUpdateRequest) (domain.Activity, error)
	Delete(Id int) error
	GetById(Id int) (domain.Activity, error)
	GetAll() ([]domain.Activity, error)
}
