package repository

import (
	database "golan-quickstart/dbconfig"
	model "golan-quickstart/models"
)

func AddUser(user model.UserInformation) (model.UserInformation, error) {
	result := database.Db.Create(&user)
	if result.Error != nil {
		return model.UserInformation{}, result.Error
	}
	return user, nil
}

func GetAllUser() []model.UserInformation {
	var users []model.UserInformation

	result := database.Db.Find(&users)
	if result.Error != nil {
		return []model.UserInformation{}
	}

	return users
}

func GetUserInfoByID(id string) []model.UserInformation {
	var user model.UserInformation
	result := database.Db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return []model.UserInformation{}
	}
	return []model.UserInformation{user}
}

func GetUserByEmail(email string) (model.UserInformation, error) {
	var user model.UserInformation
	result := database.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
