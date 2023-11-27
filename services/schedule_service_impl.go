package services

import (
	"errors"

	"github.com/agusheryanto182/go-schedule/helpers"
	"github.com/agusheryanto182/go-schedule/models/domain"
	"github.com/agusheryanto182/go-schedule/models/web"
	"github.com/agusheryanto182/go-schedule/repositories"
	"github.com/go-playground/validator/v10"
)

type ScheduleServiceImpl struct {
	ScheduleRepository repositories.ScheduleRepository
	Validate           *validator.Validate
}

func (self *ScheduleServiceImpl) Create(req web.ScheduleCreateRequest) (web.ScheduleResponse, error) {
	err := self.Validate.Struct(req)
	if err != nil {
		return web.ScheduleResponse{}, err
	}

	request := domain.Schedule{}
	request.Title = req.Title
	request.Day = req.Day
	request.UserId = req.UserId

	newSchedule := self.ScheduleRepository.Save(request)

	return helpers.ToScheduleResponse(newSchedule), nil

}

func (self *ScheduleServiceImpl) Update(req web.ScheduleUpdateRequest) (web.ScheduleResponse, error) {
	err := self.Validate.Struct(req)
	if err != nil {
		return web.ScheduleResponse{}, err
	}

	schedule, err := self.ScheduleRepository.FindById(req.ScheduleId)
	if err != nil {
		return web.ScheduleResponse{}, err
	}

	schedule.Day = req.Day
	schedule.Title = req.Title

	updated := self.ScheduleRepository.Update(schedule)

	return helpers.ToScheduleResponse(updated), nil

}

func (self *ScheduleServiceImpl) Delete(scheduleId int) error {
	err := self.ScheduleRepository.Delete(scheduleId)
	if err != nil {
		return err
	}
	return nil
}

func (self *ScheduleServiceImpl) FindAll() []web.ScheduleResponse {
	listSchedule := self.ScheduleRepository.FindAll()

	return helpers.ToScheduleResponses(listSchedule)
}

func (self *ScheduleServiceImpl) FindById(scheduleId int) (web.ScheduleResponse, error) {
	schedule, err := self.ScheduleRepository.FindById(scheduleId)
	if err != nil {
		return web.ScheduleResponse{}, err
	}

	if schedule.ScheduleId == 0 {
		return web.ScheduleResponse{}, errors.New("Schedule with that id not found")
	}

	return helpers.ToScheduleResponse(schedule), nil
}

func (self *ScheduleServiceImpl) FindByDay(day string) (web.ScheduleResponse, error) {
	schedule, err := self.ScheduleRepository.FindByDay(day)
	if err != nil {
		return web.ScheduleResponse{}, err
	}

	if schedule.ScheduleId == 0 {
		return web.ScheduleResponse{}, errors.New("Schedule is not found")
	}

	return helpers.ToScheduleResponse(schedule), nil
}

func NewScheduleService(ScheduleRepository repositories.ScheduleRepository, Validate *validator.Validate) ScheduleService {
	return &ScheduleServiceImpl{ScheduleRepository: ScheduleRepository, Validate: Validate}
}
