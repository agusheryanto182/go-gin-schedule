package domain

import "time"

type User struct {
	UserId    int `gorm:"primaryKey"`
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
