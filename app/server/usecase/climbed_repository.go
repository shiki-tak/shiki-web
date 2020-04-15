package usecase

import "github.com/shiki-tak/shiki-web/app/server/domain"

type ClimbedMountainRepository interface {
	Store(domain.ClimbedMountain) error
	Update(domain.ClimbedMountain) error
	FindById(int) (domain.ClimbedMountain, error)
	Finds() ([]domain.ClimbedMountain, error)
	Delete(int) error
}
