package router

import (
	"github.com/gin-gonic/gin"
	"wanderloge/controller"
)

func BookingRouter(router *gin.Engine) {
	router.POST("/booking", controller.CreateBooking)

}
