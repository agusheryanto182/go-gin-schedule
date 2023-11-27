package services

import "github.com/agusheryanto182/go-schedule/models/web"

type ScheduleService interface {
	FindAll() []web.ScheduleResponse
	FindById(scheduleId int) (web.ScheduleResponse, error)
	FindByDay(day string) (web.ScheduleResponse, error)
	Create(req web.ScheduleCreateRequest) (web.ScheduleResponse, error)
	Update(req web.ScheduleUpdateRequest) (web.ScheduleResponse, error)
	Delete(scheduleId int) error
}
