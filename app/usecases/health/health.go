package health

import (
	"errors"
	"time"

	"api.default.marincor.pt/app/repository/health"
	"api.default.marincor.pt/entity"
)

var errOutOfSync = errors.New("database is out of sync")

type Usecase struct {
	repo *health.Repository
}

func New() *Usecase {
	return &Usecase{
		repo: health.New(),
	}
}

func (usecase *Usecase) Check() (*entity.Health, error) {
	now := time.Now()

	if err := usecase.repo.Insert(now); err != nil {
		return nil, err
	}

	check, err := usecase.repo.GetOne(now)
	if err != nil {
		return nil, err
	}

	if check.Sync == nil || check.Sync.IsZero() {
		return nil, errOutOfSync
	}

	return check, nil
}
