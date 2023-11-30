package repositories

import "github.com/agusheryanto182/go-schedule/models/domain"

type ScheduleRepository interface {
	FindById(scheduleId int) (domain.Schedule, error)
	FindByUserId(userId int) ([]domain.Schedule, error)
	FindByDayAndUserId(day string, userId int) ([]domain.Schedule, error)
	FindByDay(day string) ([]domain.Schedule, error)
	Save(schedule domain.Schedule) domain.Schedule
	Update(schedule domain.Schedule) domain.Schedule
	Delete(scheduleId int) error
}
