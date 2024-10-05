package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID       uint `gorm:"autoIncrement:primaryKey"`
	Email    string
	Password string
}

type Event struct {
	ID          uint `gorm:"autoIncrement:primaryKey"`
	Name        string
	Description string
	Location    string
	EventDate   time.Time
	UserID      uint
	User        User
}

type Registrations struct {
	ID      uint `gorm:"autoIncrement:primaryKey"`
	UserID  uint
	EventID uint
	User    User
	Event   Event
}

func InitDB() {
	connStr := "user=deepak password=1492 dbname=PLANB host=localhost sslmode=disable"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var ver string

	db.Raw("select version()").Scan(&ver)

	fmt.Println(ver)

	err = db.AutoMigrate(&User{}, &Event{}, &Registrations{})

	if err != nil {
		panic(err)
	}

	sqlDb, _ := db.DB()

	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(5)

	DB = db
}
