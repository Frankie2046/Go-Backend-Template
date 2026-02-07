package service

import (
	"errors"

	"project/internal/model"
	"project/internal/repo"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*model.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user id")
	}
	return s.repo.GetByID(id)
}