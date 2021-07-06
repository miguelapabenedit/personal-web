package repository

import "github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/models"

type Repository interface {
	GetGratitudes() ([]models.Gratitude, error)
}
