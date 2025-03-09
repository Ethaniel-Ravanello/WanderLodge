package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wanderloge/database"
	"wanderloge/helpers"
	"wanderloge/repository"
	"wanderloge/structs"
)

func CreateListing(ctx *gin.Context) {
	var response structs.Message
	var tempListing structs.Listing

	authToken := ctx.GetHeader("Authorization")
	dataToken, err := helpers.DecodeToke(authToken)
	if err != nil {
		response = structs.Message{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: "Invalid token: " + err.Error(),
			Data:    nil,
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	err = ctx.BindJSON(&tempListing)
	if err != nil {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Invalid request body: " + err.Error(),
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tempListing.HostId = dataToken.Id

	if tempListing.Title == "" || tempListing.Description == "" || tempListing.Location == "" ||
		tempListing.Address == "" || tempListing.MaxPeople <= 0 || tempListing.PricePerNight <= 0 {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Missing required fields or invalid values",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	listingID, createdAt, approvedStatus, err := repository.CreateListing(database.DbConnection, tempListing)
	if err != nil {
		response = structs.Message{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "Failed to create listing: " + err.Error(),
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	tempListing.Id = listingID
	tempListing.CreatedAt = createdAt
	tempListing.ApprovalStatus = approvedStatus
	response = structs.Message{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Listing created successfully and sent for approval",
		Data:    tempListing,
	}
	ctx.JSON(http.StatusOK, response)
}

func GetAllListings(ctx *gin.Context) {
	var response structs.Message
	idQuery, _ := ctx.GetQuery("id")
	intId, _ := strconv.Atoi(idQuery)
	statusQuery, _ := ctx.GetQuery("approvalStatus")

	listingData, err := repository.GetListings(database.DbConnection, statusQuery, intId)

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
		Message: "Succesfuly Get All Listing",
		Data:    listingData,
	}
	ctx.JSON(http.StatusOK, response)
}

func GetListingByListId(ctx *gin.Context) {
	var response structs.Message
	listingId := ctx.Param("id")
	intId, _ := strconv.Atoi(listingId)

	listingData, err := repository.GetListingByListingId(database.DbConnection, intId)

	if err != nil {
		response = structs.Message{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "Error Getting Listing",
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response = structs.Message{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Succesfuly Get Listing",
		Data:    listingData,
	}
	ctx.JSON(http.StatusOK, response)
}

func GetListingByHostId(ctx *gin.Context) {
	var response structs.Message
	hostId := ctx.Param("id")
	intId, _ := strconv.Atoi(hostId)

	listingData, err := repository.GetListingByHostId(database.DbConnection, intId)

	if err != nil {
		response = structs.Message{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "Error Getting Listing",
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response = structs.Message{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Succesfuly Get Listing",
		Data:    listingData,
	}
	ctx.JSON(http.StatusOK, response)
}

func UpdateListing(ctx *gin.Context) {
	var response structs.Message
	var updateListing structs.Listing
	listingId := ctx.Param("id")
	intId, _ := strconv.Atoi(listingId)

	err := ctx.BindJSON(&updateListing)

	if updateListing.Title == "" || updateListing.Description == "" || updateListing.Location == "" || updateListing.Address == "" || updateListing.MaxPeople == 0 || updateListing.PricePerNight == 0 || err != nil {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Invalid or missing request body",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	err = repository.UpdateListing(database.DbConnection, updateListing, intId)

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
		Message: "Successfully Update The Listing",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, response)
}

func DeleteListing(ctx *gin.Context) {
	var response structs.Message
	listingId := ctx.Param("id")
	authToken := ctx.GetHeader("Authorization")
	dataToken, _ := helpers.DecodeToke(authToken)
	intId, err := strconv.Atoi(listingId)

	if err != nil {
		panic(err)
	}
	listingData, _ := repository.GetListingByListingId(database.DbConnection, intId)

	if dataToken.Id != listingData.HostId {
		response = structs.Message{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: "Can not delete other people listing",
			Data:    nil,
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
	fmt.Println(intId)
	err = repository.DeleteListing(database.DbConnection, intId)

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
		Error:   true,
		Message: "Succesfuly Delete Listing",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, response)
	return
}
