package user

import (
	"strings"
	db "WebService/core/database"
	userModels "WebService/apps/users/ormModels"
	userEntities "WebService/apps/users/entities"
	"encoding/json"
)

// CreateUser Create new user instance
func CreateUser(user *userEntities.User) bool{
	userOrm := &userModels.UserORM{}
	userJSON, _ := json.Marshal(&user)
	json.NewDecoder(strings.NewReader(string(userJSON))).Decode(userOrm)
	db.GetDB().Create(&userOrm)
	if userOrm.ID <= 0 {
		return false
	}
	return true
}

// ListAllUsers List all users
func ListAllUsers() (bool, []userEntities.User) {
	users := make([]*userModels.UserORM, 0)
	err := db.GetDB().Find(&users).Error
	if err != nil {
		return false, nil
	}
	return true, DecodeOrmUsers(users)
}

// DecodeOrmUser Decode ORMUser to User entity
func DecodeOrmUser(ormUser *userModels.UserORM) userEntities.User{
	user := userEntities.User {
		ID: ormUser.ID,
		Name: ormUser.Name,
		LastName: ormUser.LastName,
		Email: ormUser.Email,
		Phone: ormUser.Phone,
		Active: ormUser.Active,
	}
	return user
}

// DecodeOrmUsers Decode Array of ORMUser to User entity
func DecodeOrmUsers(ormUsers []*userModels.UserORM) []userEntities.User{
	users := []userEntities.User{}
	for _, ormUser := range ormUsers {
		users = append(users, DecodeOrmUser(ormUser))
	}
	return users
}