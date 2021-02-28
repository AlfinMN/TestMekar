package user

import (
	"BackendTestMekar/model"
	"fmt"
	"log"
)

type UserUsacaseImpl struct {
	userRepo UserRepository
}

// get all user
func (User UserUsacaseImpl) GetData() ([]*model.User, error) {

	data, err := User.userRepo.GetData()

	if err != nil {
		return nil, err
	}
	return data, nil
}

// crete user
func (User UserUsacaseImpl) CreateUser(dataUser model.User) error {
	err := User.userRepo.CreateUser(dataUser)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// get profession
func (User UserUsacaseImpl) GetProfession() ([]*model.Profession, error) {

	data, err := User.userRepo.GetProfession()

	fmt.Println("data nih", data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// get Education
func (User UserUsacaseImpl) GetEducation() ([]*model.Education, error) {

	data, err := User.userRepo.GetEducation()

	fmt.Println("data nih", data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (User UserUsacaseImpl) DeleteUser(id string) error {
	err := User.userRepo.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InitUsecaseImpl(userRepo UserRepository) UserUsecase {
	return &UserUsacaseImpl{userRepo}
}
