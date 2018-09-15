package usecase

import (
	"api/domain"
)

type UserUsecase struct {
	UserRepository UserRepository
}

func (usecase *UserUsecase) Create(u *domain.User) error {
	u.Initialize()
	err := usecase.UserRepository.Create(u)
	return err
}
