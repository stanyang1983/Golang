package service

import (
	"gin3/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userList = []pojo.User{}

// Get User
func FindAllUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, userList)
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
