package database

import (
	"encoding/json"

	"github.com/shiki-tak/shiki-web/app/server/domain"
)

type ClimbedMountainRepository struct {
	MongoDBHandler
}

func (repo *ClimbedMountainRepository) Store(m domain.ClimbedMountain) error {

	return nil
}

func (repo *ClimbedMountainRepository) FindById(key string) (domain.ClimbedMountain, error) {
	resAsBytes, err := repo.Get(key)
	if err != nil {
		return domain.ClimbedMountain{}, err
	}
	mountain := new(domain.ClimbedMountain)
	err = json.Unmarshal(resAsBytes, mountain)
	if err != nil {
		return domain.ClimbedMountain{}, err
	}

	return *mountain, nil
}

func (repo *ClimbedMountainRepository) FindAll() ([]domain.ClimbedMountain, error) {
	ｍountains := []domain.ClimbedMountain{}

	return ｍountains, nil
}
