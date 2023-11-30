package web

import "time"

type UserCreateRequest struct {
	Email string `binding:"required" validate:"required" json:"email"`
}

type UserResponse struct {
	UserId    int       `gorm:"primaryKey" json:"id"`
	Email     string    `binding:"required" validate:"required" json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
