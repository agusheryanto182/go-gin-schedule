package services

import (
	"github.com/agusheryanto182/go-todo/models/domain"
	"github.com/agusheryanto182/go-todo/models/web"
	"github.com/agusheryanto182/go-todo/repositories"
	"github.com/go-playground/validator/v10"
)

type ActivityServiceImpl struct {
	activityRepository repositories.ActivityRepository
	validate           *validator.Validate
}

func (s *ActivityServiceImpl) Create(input web.ActivityCreateRequest) (domain.Activity, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return domain.Activity{}, err
	}

	activity := domain.Activity{}
	activity.Email = input.Email
	activity.Title = input.Title

	newTodo, err := s.activityRepository.Save(activity)
	if err != nil {
		return newTodo, err
	}

	return newTodo, nil

}

func (s *ActivityServiceImpl) Update(input web.ActivityUpdateRequest) (domain.Activity, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return domain.Activity{}, err
	}

	activity, err := s.activityRepository.FindById(input.ActivityId)
	if err != nil {
		return activity, err
	}

	if input.Title == "" {
		input.Title = activity.Title
	}

	if input.Email == "" {
		input.Email = activity.Email
	}

	activity.Email = input.Email
	activity.Title = input.Title

	updated, err := s.activityRepository.Update(activity)
	if err != nil {
		return updated, err
	}

	return updated, nil

}

func (s *ActivityServiceImpl) Delete(Id int) error {
	err := s.activityRepository.Delete(Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *ActivityServiceImpl) GetById(Id int) (domain.Activity, error) {
	activity, err := s.activityRepository.FindById(Id)
	if err != nil {
		return activity, err
	}

	if activity.ActivityId <= 0 {
		return activity, err
	}

	return activity, nil
}

func (s *ActivityServiceImpl) GetAll() ([]domain.Activity, error) {
	todo, err := s.activityRepository.FindAll()
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func NewActivityService(activityRepository repositories.ActivityRepository, validate *validator.Validate) ActivityService {
	return &ActivityServiceImpl{
		activityRepository: activityRepository,
		validate:           validate,
	}
}
