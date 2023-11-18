package web

import (
	"time"
)

type ActivityCreateRequest struct {
	Email string `validate:"required,max=30" json:"email"`
	Title string `validate:"required,max=30" json:"title"`
}

type ActivityUpdateRequest struct {
	ActivityId int    `gorm:"primaryKey" validate:"required"`
	Email      string `validate:"required,max=30" json:"email"`
	Title      string `validate:"required,max=30" json:"title"`
}

type ActivityResponse struct {
	ActivityId int        `gorm:"primaryKey" json:"activity_id"`
	Email      string     `json:"email"`
	Title      string     `json:"title"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
