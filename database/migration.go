package database

import (
	"log"

	"github.com/agusheryanto182/go-schedule/models/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	log.Println("Initiating migration...")
	err := db.Migrator().AutoMigrate(
		&domain.User{},
		&domain.Schedule{},
	)
	if err != nil {
		panic(err)
	}
	log.Println("Migration Completed...")
}
