package pojo

import (
	"gin3/database"
)

//定義使用者格式
type User struct {
	Id       int    `json:Id`
	Name     string `json:Name`
	Password string `json:Password`
	Email    string `json:Email`
}

func FindAllUsers() []User {
	var users []User
	database.DBConnect.Find(&users)
	return users
}

func FindByUserID(userId string) User {
	var user User
	database.DBConnect.Where("id = ?", userId).First(&user)
	return user
}
