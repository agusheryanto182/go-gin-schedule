package domain

import (
	"database/sql"
	"time"
)

type Todo struct {
	TodoId          int `gorm:"primaryKey"`
	ActivityGroupId int
	Title           string
	IsActive        bool
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
}
