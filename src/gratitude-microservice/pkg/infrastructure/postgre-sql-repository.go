package infrastructure

import (
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/models"
	"github.com/miguelapabenedit/personal-web/src/gratitude-microservice/app/repository"
)

type repo struct{}

func NewPostgreSQL() repository.Repository {
	return &repo{}
}

func (r *repo) GetGratitudes() ([]models.Gratitude, error) {

	return nil, nil
}
