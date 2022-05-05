package src

import (
	"gin3/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	user.GET("/", service.FindAllUser)
	user.POST("/", service.PostUser)
	//delete 刪除
	//put 更改
}
