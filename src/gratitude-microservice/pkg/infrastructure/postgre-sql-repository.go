package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/models"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repo struct{}

var db *gorm.DB

var (
	server          = os.Getenv("DB_SERVER")
	port            = os.Getenv("DB_PORT")
	password        = os.Getenv("DB_PASS")
	user            = os.Getenv("DB_USER")
	database        = os.Getenv("DB_NAME")
	retryAttemps    = 0
	maxRetryAttemps = 3
)

func NewPostgreSQL() repository.Repository {

	db = openConnection()
	db.AutoMigrate(&models.Gratitude{})

	return &repo{}
}

func openConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", server, user, password, database, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		retryConnection()
	}

	return db
}

func retryConnection() {
	if retryAttemps >= maxRetryAttemps {
		panic("Retry connection attempts exceded the max stablished")
	}

	time.AfterFunc(5*time.Second, func() {
		retryAttemps++
		log.Printf("Retrying to connect to the database %v/%v", retryAttemps, maxRetryAttemps)
		openConnection()
	})

}

func (r *repo) GetGratitudes() ([]models.Gratitude, error) {
	var gratitudes = []models.Gratitude{}

	result := db.Find(&gratitudes)

	if result.Error != nil {
		return nil, result.Error
	}

	return gratitudes, nil
}
