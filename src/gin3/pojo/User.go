package pojo

import (
	"gin3/database"
)

//定義使用者格式
type User struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name" binding:"required,gt=5"`
	Password string `json:"Password" binding:"required,min=4,max=20,userpassword"`
	Email    string `json:"Email" binding:"required,email"`
}
type Users struct {
	UserList     []User `json:"UserList" binding:"required,gt=0,lt=3"`
	UserListSize int    `json:"UserListSize"`
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

//Post
func CreateUser(user User) User {
	database.DBConnect.Create(&user)
	return user
}

//Delete
func DeleteUser(userId string) bool {
	user := User{}
	result := database.DBConnect.Where("id = ?", userId).Delete(&user)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

//Update //目前看不懂...先用不到
func UpdateUser(userId string, user User) User {
	database.DBConnect.Model(&user).Where("id = ?", userId).Updates(user)
	return user
}
