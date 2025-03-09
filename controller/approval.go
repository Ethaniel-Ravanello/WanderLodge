package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wanderloge/database"
	"wanderloge/repository"
	"wanderloge/structs"
)

func ActionApproval(ctx *gin.Context) {
	var response structs.Message

	// Parse approvalId from query parameters
	approvalId := ctx.Param("approvalId")
	intApprovalId, err := strconv.Atoi(approvalId)
	if err != nil {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Invalid approval ID",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Parse newStatus from request body
	var requestBody struct {
		NewStatus string `json:"newStatus"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Invalid request body",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate new status
	if requestBody.NewStatus == "" {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Status cannot be empty",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Call repository function to update approval status
	err = repository.ActionApproval(database.DbConnection, intApprovalId, requestBody.NewStatus)
	if err != nil {
		response = structs.Message{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "Failed to update approval status",
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	// Success response
	response = structs.Message{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Approval status updated successfully",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, response)
	return
}
