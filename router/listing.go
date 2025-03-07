package router

import (
	"github.com/gin-gonic/gin"
	"wanderloge/controller"
)

func ListingRouter(router *gin.Engine) {
	router.POST("/listing", controller.CreateListing)
	router.GET("/listing", controller.GetAllListings)
	router.GET("/listing/id/:id", controller.GetListingByListId)
	router.GET("/listing/hostId/:id", controller.GetListingByHostId)
	router.PUT("/listing/id/:id", controller.UpdateListing)
	router.DELETE("/listing/:id", controller.DeleteListing)
}
