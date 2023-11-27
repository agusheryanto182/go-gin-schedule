package domain

import "time"

type Schedule struct {
	ScheduleId int `gorm:"primaryKey"`
	UserId     int
	Title      string
	Day        string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
