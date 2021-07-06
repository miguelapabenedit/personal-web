package services

import "github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/models"

type Service interface {
	GetGratitudes() ([]models.Gratitude, error)
}
