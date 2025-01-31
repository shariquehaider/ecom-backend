package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shariquehaider/ecom-backend/models"
	"github.com/shariquehaider/ecom-backend/utils"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type updatePasswordRequest struct {
	Password           string `json:"currentPassword"`
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

func LoginController(ctx *gin.Context) {
	var loginCredentials Login

	if err := ctx.ShouldBindJSON(&loginCredentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := models.FindUserByUsername(loginCredentials.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	isVerified := utils.CompareHashPassword(loginCredentials.Password, user.Password)
	if !isVerified {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateJWT(user.Id.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func RegisterController(ctx *gin.Context) {
	var registerCreds models.User

	if err := ctx.ShouldBindJSON(&registerCreds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, _ := models.FindUserByUsername(registerCreds.Email)
	if user != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "User Already Exist!",
		})
		return
	}

	hashedPassword, err := utils.GenerateHashPassword(registerCreds.Password)
	if err != nil {
		panic(err)
	}

	newUser := models.User{
		Email:    registerCreds.Email,
		Username: registerCreds.Username,
		Password: hashedPassword,
		Name:     registerCreds.Name,
	}

	err = models.CreateUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully. Login to Continue."})
}

func GetProfileController(ctx *gin.Context) {
	userID := ctx.MustGet("_id").(string)
	user, err := models.FindById(userID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func ChangePasswordController(ctx *gin.Context) {
	userID := ctx.MustGet("_id").(string)
	var newPassword updatePasswordRequest

	if err := ctx.ShouldBindJSON(&newPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	fmt.Println(newPassword)

	isMatched := utils.VerifyNewPassword(newPassword.NewPassword, newPassword.ConfirmNewPassword)
	if !isMatched {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "New Password MisMatched"})
		return
	}

	_, err := models.ChangePasswordByID(userID, newPassword.Password, newPassword.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "Password Changed!"})

}
