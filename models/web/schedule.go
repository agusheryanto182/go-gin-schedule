package web

import "time"

type ScheduleCreateRequest struct {
	UserId int    `validate:"required" json:"user_id"`
	Title  string `binding:"required" validate:"required,max=30" json:"title"`
	Day    string `binding:"required" validate:"required,max=30" json:"day"`
}

type ScheduleUpdateRequest struct {
	ScheduleId int    `gorm:"primaryKey" validate:"required,max=30" json:"schedule_id"`
	UserId     int    `validate:"required,max=30" json:"user_id"`
	Title      string `binding:"required" validate:"required,max=30" json:"title"`
	Day        string `binding:"required" validate:"required,max=30" json:"day"`
}

type ScheduleResponse struct {
	ScheduleId int       `gorm:"primaryKey" json:"schedule_id"`
	UserId     int       `json:"user_id"`
	Title      string    `json:"title"`
	Day        string    `json:"day"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
