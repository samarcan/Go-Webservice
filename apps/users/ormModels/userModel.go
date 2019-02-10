package users

import (
	"github.com/jinzhu/gorm"
)

// UserORM Struct of our User model
type UserORM struct {
	gorm.Model
	Name string
	LastName string
	Email string
	Phone string
	Active bool
}
// TableName Change table name
func (UserORM) TableName() string {
	return "users"
  }