package user

import "BackendTestMekar/model"

type UserUsecase interface {
	GetData() ([]*model.User, error)
	CreateUser(user model.User) error
	DeleteUser(string) error
	GetProfession() ([]*model.Profession, error)
	GetEducation() ([]*model.Education, error)
}
