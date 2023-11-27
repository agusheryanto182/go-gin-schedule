package repositories

import "github.com/agusheryanto182/go-schedule/models/domain"

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	FindById(userId int) (domain.User, error)
}
