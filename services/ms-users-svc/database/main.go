package database

import (
	"fmt"

	"github.com/mandaputtra/projectsprint-projects2/services/ms-users-svc/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Preferences string
	WeightUnit  string
	HeightUnit  string
	Weight      float64
	Height      float64
	Name        string
	ImageUri    string
}

// Setup database
var db *gorm.DB

func ConnectDatabase(env config.Environment) *gorm.DB {
	fmt.Println("Connect to database ....")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s",
		env.DATABASE_HOST,
		env.DATABASE_USER,
		env.DATABASE_PASSWORD,
		env.DATABASE_NAME,
		env.DATABASE_PORT,
		env.DATABASE_SCHEMA,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Test query
	var result string
	db.Raw("SELECT 1;").Scan(&result)

	fmt.Printf("Connection successfull. Result from test SQL: %s\n", result)

	// Migrations
	db.AutoMigrate(&User{})
	return db
}

func GetDB() *gorm.DB {
	return db
}
