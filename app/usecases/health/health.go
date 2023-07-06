package health

import (
	constants "api.default.marincor/app/errors"
	"api.default.marincor/app/repository/health"
	"api.default.marincor/entity"
)

type Usecase struct {
	repo *health.Repository
}

func New() *Usecase {
	return &Usecase{
		repo: health.New(),
	}
}

func (uc *Usecase) Check() (*entity.Health, error) {
	testDatabase, err := uc.repo.GetHealthCheck()
	if err != nil {
		panic(constants.ErrDatabaseNotConnected)
	}

	return testDatabase, nil
}
