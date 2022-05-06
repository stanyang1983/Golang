package src

import (
	"gin3/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	user.GET("/", service.FindAllUser)
	user.GET("/:id", service.FindByUserId)
	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUserList)
	user.DELETE("/:id", service.DeleteUser)
	user.PUT("/:id", service.UpdateUser)
}
