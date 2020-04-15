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

func (interactor *ClimbedMountainInteractor) Update(m domain.ClimbedMountain) error {
	return interactor.ClimbedRepository.Update(m)
}

func (interactor *ClimbedMountainInteractor) GetById(id int) (domain.ClimbedMountain, error) {
	climbedMountain, err := interactor.ClimbedRepository.FindById(id)
	if err != nil {
		return domain.ClimbedMountain{}, err
	}
	return climbedMountain, nil
}

func (interactor *ClimbedMountainInteractor) Gets() ([]domain.ClimbedMountain, error) {
	climbedMountains, err := interactor.ClimbedRepository.Finds()
	if err != nil {
		return nil, err
	}
	return climbedMountains, nil
}

func (interactor *ClimbedMountainInteractor) Remove(id int) error {
	err := interactor.ClimbedRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
