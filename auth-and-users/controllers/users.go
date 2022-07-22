package controllers

import (
	"auth-and-users/models"
	"auth-and-users/repository"
	"auth-and-users/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := repository.DB.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}

func WhoAmI(context *gin.Context) {
	tokenString := context.GetHeader("Authorization")
	var user models.User
	if tokenString == "" {
		context.JSON(401, gin.H{"error": "request does not contain an access token"})
		context.Abort()
		return
	}
	uid, err := utils.DecodeToken(tokenString)
	if err != nil {
		context.JSON(401, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := repository.DB.Where("id = ?", uid).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"username": user.Name, "id": user.ID, "email": user.Email, "createdAt": user.CreatedAt})

}
