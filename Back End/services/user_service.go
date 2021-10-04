package services

import (
	"webTech/models/domains"
	"webTech/repositories"
)

type IUserService interface {
	CreateUser(user *domains.User) (*domains.User, error)
}

type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(
	userRepository repositories.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (that *UserService) CreateUser(user *domains.User) (*domains.User, error) {
	//logical or business op
	createdUser, err := that.userRepository.Create(user)
	return nil, nil
}
