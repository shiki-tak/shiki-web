package usecase

import "github.com/shiki-tak/shiki-web/app/server/domain"

type ClimbedMountainRepository interface {
	Store(domain.ClimbedMountain) error
	FindById(string) (domain.ClimbedMountain, error)
	FindAll() ([]domain.ClimbedMountain, error)
}
