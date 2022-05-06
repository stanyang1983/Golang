package service

import (
	"gin3/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//var userList = []pojo.User{}

// Get User
func FindAllUser(ctx *gin.Context) {
	//ctx.JSON(http.StatusOK, userList)
	users := pojo.FindAllUsers()
	ctx.JSON(http.StatusOK, users)
}

//Get User by id

func FindByUserId(ctx *gin.Context) {

	user := pojo.FindByUserID(ctx.Param("id"))
	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, "Not Found")
		return
	}
	log.Println("User ->", user)
	ctx.JSON(http.StatusOK, user)

}

// Post User
func PostUser(ctx *gin.Context) {
	user := pojo.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.String(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	newUser := pojo.CreateUser(user)
	ctx.JSON(http.StatusCreated, newUser)

}

// Delete User
func DeleteUser(ctx *gin.Context) {
	user := pojo.DeleteUser(ctx.Param("id"))
	if !user {
		ctx.JSON(http.StatusNotFound, "Not Found")
		return
	}
	ctx.JSON(http.StatusOK, "Delete Sussessful")
}

// Update User
func UpdateUser(ctx *gin.Context) {
	user := pojo.User{}
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = pojo.UpdateUser(ctx.Param("id"), user)
	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, "Not Found")
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// CreateUserList
//目前看不懂...先用不到
func CreateUserList(ctx *gin.Context) {
	users := pojo.Users{}
	err := ctx.BindJSON(&users)
	if err != nil {
		ctx.String(400, "Error:%s", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, users)
}
