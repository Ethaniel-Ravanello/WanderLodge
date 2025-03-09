package router

import (
	"github.com/gin-gonic/gin"
	"wanderloge/controller"
)

func AprrovalRouter(router *gin.Engine) {
	router.PUT("/approval/:approvalId", controller.ActionApproval)

}
