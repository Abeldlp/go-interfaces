package service

import "github.com/Abeldlp/interfaces/model"

type UserService struct {
	Users []model.User
}

func NewUserService() Service[model.User] {
	return &UserService{}
}

func (u *UserService) GetAll() []model.User {
	return u.Users
}

func (u *UserService) Create(user model.User) model.User {
	user.Id = len(u.Users) + 1
	u.Users = append(u.Users, user)
	return user
}
