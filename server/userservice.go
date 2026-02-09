package service

import (
	"errors"

	model "golan-quickstart/models"
	userrepo "golan-quickstart/repository"
	utils "golan-quickstart/utils"
)

func UserVerification(user model.UserInformation) (model.UserInformation, error) {
	if user.Email == "" {
		return model.UserInformation{}, errors.New("email cannot be empty")
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		return model.UserInformation{}, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	savedUser, err := userrepo.AddUser(user)
	if err != nil {
		return model.UserInformation{}, err
	}

	return savedUser, nil
}

func GetAllUser() []model.UserInformation {
	return userrepo.GetAllUser()
}

func GetUserById(id string) []model.UserInformation {
	if id == " " {
		return []model.UserInformation{}
	}
	return userrepo.GetUserInfoByID(id)
}
