package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	UserRouter(router)
	ListingRouter(router)
	BookingRouter(router)
	AprrovalRouter(router)

	return router
}
