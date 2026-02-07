package repo

import "project/internal/model"

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetByID(id int) (*model.User, error) {
	// mock 数据
	return &model.User{
		ID:   id,
		Name: "fiber-user",
	}, nil
}