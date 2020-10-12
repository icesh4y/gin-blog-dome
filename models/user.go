package models

// 用户

import (
	"essential/dao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Phone string `gorm:"unique;not null"`
	Pwd   string `gorm:"size(255);not null"`
}

func SelectUserPhone(phone string) bool {
	var user User
	dao.DB.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
func CreateNewUser(name, phone, pwd string) {
	newUser := User{
		Name:  name,
		Phone: phone,
		Pwd:   pwd,
	}
	dao.DB.Debug().Create(&newUser)
}

func SelectUser(phone string) *User {
	var user User
	dao.DB.Debug().Where("phone = ?", phone).First(&user)
	return &user
}
