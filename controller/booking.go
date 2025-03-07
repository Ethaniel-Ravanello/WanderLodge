package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"wanderloge/database"
	"wanderloge/repository"
	"wanderloge/structs"
)

func GetAllBooking(ctx *gin.Context) {
	var response structs.Message
	var userBooking *[]structs.Booking
	userId := ctx.Query("guestId")
	intId, _ := strconv.Atoi(userId)
	userType := ctx.Query("userType")

	if userType == "host" {
		booking, _ := repository.GetAllBookingForHost(database.DbConnection, intId)
		userBooking = &booking
	} else if userType == "guest" {
		booking, _ := repository.GetAllBookingForGuest(database.DbConnection, intId)
		userBooking = &booking
	} else {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Invalid User Type",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response = structs.Message{
		Code:    http.StatusOK,
		Error:   true,
		Message: "Successfully get all Booking Data",
		Data:    userBooking,
	}
	ctx.JSON(http.StatusOK, response)
	return
}

func CreateBooking(ctx *gin.Context) {
	var response structs.Message
	var tempBooking structs.Booking
	nullApproverId := sql.NullInt32{Valid: false}

	err := ctx.BindJSON(&tempBooking)

	if err != nil {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Error Getting Body",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err = repository.AddBookListing(database.DbConnection, tempBooking)
	if err != nil {
		response = structs.Message{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	err = repository.CreateApproval(database.DbConnection, tempBooking.ListingId, "booking", nullApproverId, "pending", time.Now(), time.Now())
	if err != nil {
		response = structs.Message{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response = structs.Message{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Succesfully Created A Booking Order",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, response)
}
