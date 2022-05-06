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
	//delete 刪除
	user.DELETE("/:id", service.DeleteUser)
	//put 更改
	user.PUT("/:id", service.PutUser)
}
