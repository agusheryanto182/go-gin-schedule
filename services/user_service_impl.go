package services

import (
	"github.com/agusheryanto182/go-schedule/helpers"
	"github.com/agusheryanto182/go-schedule/models/domain"
	"github.com/agusheryanto182/go-schedule/models/web"
	"github.com/agusheryanto182/go-schedule/repositories"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	Validate       *validator.Validate
}

func (self *UserServiceImpl) Create(req web.UserCreateRequest) (web.UserResponse, error) {
	err := self.Validate.Struct(req)
	if err != nil {
		return web.UserResponse{}, err
	}

	request := domain.User{}
	request.Email = req.Email

	newRequest, err := self.UserRepository.Save(request)
	if err != nil {
		return web.UserResponse(newRequest), err
	}

	return helpers.ToUserResponse(newRequest), nil
}

func (self *UserServiceImpl) FindById(userId int) (web.UserResponse, error) {
	user, err := self.UserRepository.FindById(userId)
	if err != nil {
		return web.UserResponse{}, err
	}

	if user.UserId == 0 {
		return web.UserResponse{}, err
	}

	return helpers.ToUserResponse(user), nil

}

func NewUserService(UserRepository repositories.UserRepository, Validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: UserRepository, Validate: Validate}
}
