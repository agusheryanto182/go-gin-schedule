package web

import "time"

type ScheduleCreateRequest struct {
	UserId int    `validate:"required" json:"user_id"`
	Title  string `binding:"required" validate:"required" json:"title"`
	Day    string `binding:"required" validate:"required" json:"day"`
}

type ScheduleUpdateRequest struct {
	ScheduleId int    `gorm:"primaryKey" validate:"required" json:"id"`
	UserId     int    `validate:"required" json:"user_id"`
	Title      string `binding:"required" validate:"required" json:"title"`
}

type ScheduleResponse struct {
	ScheduleId int       `gorm:"primaryKey" json:"id"`
	UserId     int       `json:"user_id"`
	Title      string    `json:"title"`
	Day        string    `json:"day"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
