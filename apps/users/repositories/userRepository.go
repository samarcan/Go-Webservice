package users

import (
	userEntities "WebService/apps/users/entities"
	userDbRepository "WebService/apps/users/dbRepositories"
)

func ListAllUsers() (bool, []userEntities.User) {
	result, users := userDbRepository.ListAllUsers()
	return result, users
}

func CreateUser(user *userEntities.User) bool{
	result := userDbRepository.CreateUser(user)
	return result
}