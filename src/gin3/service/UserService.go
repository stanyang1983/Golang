package service

import (
	"gin3/pojo"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userList = []pojo.User{}

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
		ctx.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	userList = append(userList, user)
	ctx.JSON(http.StatusOK, "Sussessfully posted")

}

func DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id")) //string to int
	for _, user := range userList {
		log.Println(user)

		userList = append(userList[:userId], userList[userId+1:]...)
		ctx.JSON(http.StatusOK, "Sussessfully deleted")
		return

	}
	ctx.JSON(http.StatusNotFound, "Error")

}

func PutUser(ctx *gin.Context) {
	beforeUser := pojo.User{}
	err := ctx.BindJSON(&beforeUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Error")
	}
	userId, _ := strconv.Atoi(ctx.Param("id"))
	for key, user := range userList {
		if userId == user.Id {
			userList[key] = beforeUser
			log.Println(userList[key])
			ctx.JSON(http.StatusOK, "Sussessfully")
			return
		}
	}
	ctx.JSON(http.StatusNotFound, "Error")
}
