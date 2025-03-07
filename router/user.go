package router

import (
	"github.com/gin-gonic/gin"
	"wanderloge/controller"
)

func UserRouter(router *gin.Engine) {
	router.POST("/signup", controller.Signup)
	router.PUT("/signin", controller.Signin)
	router.GET("/user", controller.GetAllUsers)
	router.GET("/user/:id", controller.GetUserById)
	router.PUT("/user/:id", controller.UpdateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
}
