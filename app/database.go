package app

import (
	"fmt"
	"log"
	"time"

	"github.com/agusheryanto182/go-schedule/models/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(c *config.Global) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name)

	fmt.Printf("Connecting to database with DSN: %s\n", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		for i := 0; i < 10; i++ {
			fmt.Println("Retrying to connect to the database...")
			time.Sleep(5 * time.Second) // Tunggu 5 detik sebelum mencoba lagi
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err == nil {
				break
			}
		}
		if err != nil {
			log.Fatal("Gagal terkoneksi ke database setelah beberapa percobaan")
		}
	}
	fmt.Println("Connected to the database")

	return db
}
