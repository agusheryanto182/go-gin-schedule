package services

import "github.com/agusheryanto182/go-schedule/models/web"

type UserService interface {
	Create(req web.UserCreateRequest) (web.UserResponse, error)
	FindById(userId int) (web.UserResponse, error)
	FindByEmail(email string) (web.UserResponse, error)
}
