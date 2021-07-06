package services

import "github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/models"

type service struct{}

func NewGratitudeService() Service {
	return &service{}
}

func (s *service) GetGratitudes() ([]models.Gratitude, error) {
	gratitudes := []models.Gratitude{
		{ID: 1, Name: "pedro", Description: "pedro salva"},
		{ID: 2, Name: "pedro", Description: "pedro salva"},
	}

	return gratitudes, nil
}
