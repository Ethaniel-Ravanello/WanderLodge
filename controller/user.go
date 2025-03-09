package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"wanderloge/database"
	"wanderloge/helpers"
	"wanderloge/repository"
	"wanderloge/structs"
)

func Signup(ctx *gin.Context) {
	var newUser structs.User
	var response structs.Message
	err := ctx.BindJSON(&newUser)

	if err != nil {
		panic(err)
	}
	if newUser.FirstName == "" || newUser.LastName == "" || newUser.Email == "" || newUser.Password == "" || newUser.PhoneNumber == 0 || newUser.Roles == "" {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Invalid Request",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	hashPassword, _ := helpers.HashPassword(newUser.Password)
	newUser.Password = hashPassword
	fmt.Println(newUser)
	listingId, err := repository.SignUp(database.DbConnection, newUser)

	if err != nil {
		log.Fatal(err)
	}

	newUser.Id = listingId

	response = structs.Message{
		Code:    http.StatusOK,
		Error:   true,
		Message: "Success Create User",
		Data:    newUser,
	}
	ctx.JSON(http.StatusOK, response)
}

func Signin(ctx *gin.Context) {
	var loginUser structs.User
	var response structs.Message
	err := ctx.BindJSON(&loginUser)

	if err != nil {
		panic(err)
	}
	userData := repository.SignIn(database.DbConnection, loginUser.Email)
	isPassValid := helpers.CheckPasswordHash(loginUser.Password, userData.Password)
	fmt.Println(isPassValid)

	if !isPassValid {
		response = structs.Message{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: "Invalid Credentials",
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
	newToken, err := helpers.GenerateToken(loginUser.FirstName, userData.Id, userData.Roles)

	if err != nil {
		response = structs.Message{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "Error Signing Token",
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response = structs.Message{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Success Login",
		Data:    gin.H{"token": newToken},
	}
	ctx.JSON(http.StatusOK, response)
}

func GetAllUsers(ctx *gin.Context) {
	var response structs.Message
	allUsers, err := repository.GetUsers(database.DbConnection)

	if err != nil {
		response = structs.Message{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "Error Fetching All Users",
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response = structs.Message{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Success Fetching All Users",
		Data:    allUsers,
	}
	ctx.JSON(http.StatusOK, response)
}

func GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	var response structs.Message
	intUserId, _ := strconv.Atoi(userId)
	onlyUser, err := repository.GetUserById(database.DbConnection, intUserId, "")

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
		Message: "Success Fetching User",
		Data:    onlyUser,
	}
	ctx.JSON(http.StatusOK, response)
}

func UpdateUser(ctx *gin.Context) {
	var tempUser structs.User
	var response structs.Message
	userId := ctx.Param("id")
	intUserId, _ := strconv.Atoi(userId)

	err := ctx.BindJSON(&tempUser)
	tempUser.Id = intUserId
	if err != nil {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if tempUser.FirstName == "" || tempUser.LastName == "" || tempUser.Email == "" || tempUser.Password == "" || tempUser.PhoneNumber == 0 || tempUser.Roles == "" {
		response = structs.Message{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Invalid Request",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = repository.UpdateUser(database.DbConnection, tempUser, intUserId)
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
		Message: "Success Updating User",
		Data:    tempUser,
	}
	ctx.JSON(http.StatusOK, response)
}

func DeleteUser(ctx *gin.Context) {
	var response structs.Message
	authToken := ctx.GetHeader("Authorization")
	userId := ctx.Param("id")
	intId, _ := strconv.Atoi(userId)
	tokenData, err := helpers.DecodeToke(authToken)

	if err != nil {
		response = structs.Message{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: err.Error(),
			Data:    nil,
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	if tokenData.Role != "admin" {
		response = structs.Message{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: "Unauthorized Request",
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}
	err = repository.DeleteUser(database.DbConnection, intId)
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
		Message: "Success Deleting User",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, response)
}
