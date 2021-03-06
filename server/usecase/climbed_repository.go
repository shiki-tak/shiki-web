package usecase

import "github.com/shiki-tak/shiki-web/server/domain"

type ClimbedMountainRepository interface {
	Store(domain.ClimbedMountain) error
	Update(domain.ClimbedMountain) error
	FindById(int) (domain.ClimbedMountain, error)
	FindAll() ([]domain.ClimbedMountain, error)
	Delete(int) error
}
