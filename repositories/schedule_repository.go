package repositories

import "github.com/agusheryanto182/go-schedule/models/domain"

type ScheduleRepository interface {
	FindAll() []domain.Schedule
	FindById(scheduleId int) (domain.Schedule, error)
	FindByDay(day string) (domain.Schedule, error)
	Save(schedule domain.Schedule) domain.Schedule
	Update(schedule domain.Schedule) domain.Schedule
	Delete(scheduleId int) error
}
