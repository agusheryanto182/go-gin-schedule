package repositories

import (
	"github.com/agusheryanto182/go-schedule/models/domain"
	"gorm.io/gorm"
)

type ScheduleRepositoryImpl struct {
	db *gorm.DB
}

func (self *ScheduleRepositoryImpl) Save(schedule domain.Schedule) domain.Schedule {
	err := self.db.Create(&schedule).Error
	if err != nil {
		return schedule
	}
	return schedule
}

func (self *ScheduleRepositoryImpl) FindById(scheduleId int) (domain.Schedule, error) {
	var schedule domain.Schedule
	err := self.db.Where("schedule_id = ?", scheduleId).Find(&schedule).Error
	if err != nil {
		return schedule, err
	}

	return schedule, nil
}

func (self *ScheduleRepositoryImpl) FindByDay(day string) (domain.Schedule, error) {
	var schedule domain.Schedule
	err := self.db.Where("day = ?", day).Find(&schedule).Error
	if err != nil {
		return schedule, err
	}

	return schedule, nil
}

func (self *ScheduleRepositoryImpl) FindAll() []domain.Schedule {
	var listSchedule []domain.Schedule
	err := self.db.Find(&listSchedule).Error
	if err != nil {
		return listSchedule
	}
	return listSchedule
}

func (self *ScheduleRepositoryImpl) Update(schedule domain.Schedule) domain.Schedule {
	err := self.db.Save(&schedule).Error
	if err != nil {
		return schedule
	}
	return schedule
}

func (self *ScheduleRepositoryImpl) Delete(scheduleId int) error {
	err := self.db.Delete(&domain.Schedule{}, "schedule_id = ?", scheduleId).Error
	if err != nil {
		return err
	}
	return nil
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &ScheduleRepositoryImpl{db: db}
}
