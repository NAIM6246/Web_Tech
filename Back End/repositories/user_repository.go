package repositories

import "webTech/models/domains"

type IUserRepository interface {
	Create(user *domains.User) (*domains.User, error)
}
type UserRepository struct {
	db
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (that *UserRepository) Create(user *domains.User) (*domains.User, error) {

	return nil, nil
}
