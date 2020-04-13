package usecase

import (
	"github.com/shiki-tak/shiki-web/app/server/domain"
)

type ClimbedMountainInteractor struct {
	ClimbedRepository ClimbedMountainRepository
}

func (interactor *ClimbedMountainInteractor) Add(m domain.ClimbedMountain) error {
	return interactor.ClimbedRepository.Store(m)
}

func (interactor *ClimbedMountainInteractor) AllClimbedMountains() ([]domain.ClimbedMountain, error) {
	climbedMountains, err := interactor.ClimbedRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return climbedMountains, nil
}

func (interactor *ClimbedMountainInteractor) ClimbedMountainById(key string) (domain.ClimbedMountain, error) {
	climbedMountain, err := interactor.ClimbedRepository.FindById(key)
	if err != nil {
		return domain.ClimbedMountain{}, err
	}
	return climbedMountain, nil
}
