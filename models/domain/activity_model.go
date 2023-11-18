package domain

import (
	"database/sql"
	"time"
)

type Activity struct {
	ActivityId int `gorm:"primaryKey"`
	Email      string
	Title      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
}
