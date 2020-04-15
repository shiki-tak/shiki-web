package usecase

import "github.com/shiki-tak/shiki-web/app/server/domain"

type ClimbedMountainRepository interface {
	Store(domain.ClimbedMountain) error
	FindById(int) (domain.ClimbedMountain, error)
	Gets() ([]domain.ClimbedMountain, error)
}
