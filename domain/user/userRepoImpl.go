package user

import (
	"BackendTestMekar/model"
	"BackendTestMekar/utils"
	"database/sql"
)

type UserRepoImpl struct {
	DB *sql.DB
}

// get all data
func (user UserRepoImpl) GetData() ([]*model.User, error) {
	dataUsers := []*model.User{}

	data, err := user.DB.Query(utils.GET_USERS)

	if err != nil {
		return nil, err
	}

	for data.Next() {
		users := model.User{}
		err := data.Scan(&users.UserId, &users.IdentityNumber, &users.Username, &users.DateOfBirth,
			&users.Education.EducationId, &users.Education.EducationName,
			&users.Profession.ProfessionId, &users.Profession.ProfessionName, &users.UserStatus)

		if err != nil {
			return nil, err
		}

		dataUsers = append(dataUsers, &users)

	}
	return dataUsers, nil
}

// create user
func (user UserRepoImpl) CreateUser(dataUser model.User) error {

	tx, err := user.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.ADD_USER)
	defer stmt.Close()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = stmt.Exec(&dataUser.UserId, &dataUser.Username, &dataUser.IdentityNumber, &dataUser.DateOfBirth, &dataUser.Profession.ProfessionId, &dataUser.Education.EducationId)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

// get profession
func (user UserRepoImpl) GetProfession() ([]*model.Profession, error) {
	listProfession := []*model.Profession{}

	data, err := user.DB.Query(utils.GET_PROFESSION)

	if err != nil {
		return nil, err
	}
	for data.Next() {
		profession := model.Profession{}
		err := data.Scan(&profession.ProfessionId, &profession.ProfessionName)
		if err != nil {
			return nil, err
		}

		listProfession = append(listProfession, &profession)
	}
	return listProfession, nil
}

// get Education
func (user UserRepoImpl) GetEducation() ([]*model.Education, error) {
	listEducation := []*model.Education{}

	data, err := user.DB.Query(utils.GET_EDUCATION)

	if err != nil {
		return nil, err
	}
	for data.Next() {
		Education := model.Education{}
		err := data.Scan(&Education.EducationId, &Education.EducationName)
		if err != nil {
			return nil, err
		}

		listEducation = append(listEducation, &Education)
	}
	return listEducation, nil
}

func (user UserRepoImpl) DeleteUser(id string) error {
	tx, err := user.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.DELETE_USER)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func InitUserRepoImpl(DB *sql.DB) UserRepository {
	return &UserRepoImpl{DB: DB}
}
