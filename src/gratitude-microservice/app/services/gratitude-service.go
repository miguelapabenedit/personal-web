package services

import (
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/models"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/repository"
)

type service struct{}

var repo repository.Repository

func NewGratitudeService(repository repository.Repository) Service {
	repo = repository
	return &service{}
}

func (s *service) GetGratitudes() ([]models.Gratitude, error) {

	gratitudes, err := repo.GetGratitudes()

	if err != nil {
		return nil, err
	}

	return gratitudes, nil
}
