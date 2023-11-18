package web

import (
	"time"
)

type TodoCreateRequest struct {
	ActivityGroupId int    `validate:"required" json:"activity_group_id"`
	Title           string `validate:"required,max=30" json:"title"`
	Priority        string `json:"priority"`
	IsActive        bool   `json:"is_active"`
}

type TodoUpdateRequest struct {
	TodoId          int    `gorm:"primaryKey" validate:"required"`
	ActivityGroupId int    `json:"activity_group_id"`
	Title           string `json:"title"`
	Priority        string `json:"priority"`
	IsActive        *bool  `json:"is_active"`
}

type TodoResponse struct {
	TodoId          int        `gorm:"primaryKey" json:"todo_id"`
	ActivityGroupId int        `json:"activity_group_id"`
	Title           string     `json:"title"`
	IsActive        bool       `json:"is_active"`
	Priority        string     `json:"priority"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}
