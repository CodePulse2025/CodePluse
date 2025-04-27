package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/hiteshwagh9383/CodePulse/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // postgres driver
)

var db *gorm.DB

func InitDB() {
	var err error

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate (create tables automatically)
	db.AutoMigrate(&model.Task{}, &model.Developer{}, &model.Commit{}, &model.PR{}, &model.Review{})

	log.Println("Database connected & migrated successfully")

}

func GetDB() *gorm.DB {
	return db
}
