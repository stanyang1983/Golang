package main

import (
	"gin3/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	src.AddUserRouter(v1)
	// router.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message":  "ping",
	// 		"message2": "Success!",
	// 		"status":   1,
	// 	})
	// })

	// router.POST("/ping/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")
	// 	ctx.JSON(200, gin.H{
	// 		"id": id,
	// 	})
	// })
	router.Run(":8000")
}

// 使用pojo/User的規則----> service/Userservice 定義 get post  ----->  綁到 src/UserRouter  ---->引用回來 main.go

// 物件-->服務-->路由
