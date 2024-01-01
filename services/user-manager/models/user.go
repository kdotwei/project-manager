package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Roles    []Role `gorm:"many2many:user_roles;" json:"roles"`
}

type Role struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `gorm:"many2many:user_roles;" json:"-"`
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func UserHasRole(db *gorm.DB, user *User, roleName string) bool {
	var roles []Role
	db.Model(user).Association("Roles").Find(&roles)
	for _, role := range roles {
		if role.Name == roleName {
			return true
		}
	}
	return false
}

func AddRoleToUser(db *gorm.DB, user *User, role *Role) error {
	return db.Model(user).Association("Roles").Append(role)
}
