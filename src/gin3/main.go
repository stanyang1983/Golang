package main

import (
	"gin3/database"
	"gin3/middlewares"
	"gin3/pojo"
	"gin3/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func setupLogging() { //日誌功能
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	//日誌功能 setupLogging()

	router := gin.Default()

	// 註冊 Validator Func
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpassword", middlewares.UserPassword)
		v.RegisterStructValidation(middlewares.UserList, pojo.Users{}) //目前看不懂...先用不到
	}
	router.Use(gin.Recovery())

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

	go func() {
		database.DB()
	}()

	router.Run(":8000")
}

// 使用pojo/User的規則----> service/Userservice 定義 get post  ----->  綁到 src/UserRouter ----> 呼叫DBConnect ---->引用回來 main.go

// 物件-->服務-->路由
