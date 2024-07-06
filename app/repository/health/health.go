package health

import (
	"time"

	"api.default.marincor.pt/adapters/database"
	"api.default.marincor.pt/app/appinstance"
	"api.default.marincor.pt/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct {
	database *database.Database[entity.Health]
}

func New() *Repository {
	return &Repository{
		database: database.New[entity.Health](appinstance.Data.DB),
	}
}

func (repo *Repository) Insert(now time.Time) error {
	_, err := repo.database.InsertOne("health", entity.Health{
		Sync: &now,
	})

	return err
}

func (repo *Repository) GetOne(now time.Time) (*entity.Health, error) {
	return repo.database.GetOne("health", bson.M{"sync": now})
}

func (repo *Repository) ListAll() ([]*entity.Health, error) {
	return repo.database.ListAll("health")
}
