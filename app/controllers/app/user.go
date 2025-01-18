package app

import (
	"chatbot-backend/app/common/request"
	"chatbot-backend/app/common/response"
	"chatbot-backend/app/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login use google token to log in
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if user, err := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(user)
		if err != nil {
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

// Logout log out the account
func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "failed to logout")
		return
	}
	response.Success(c, nil)
}

// Info get user info
func Info(c *gin.Context) {
	if user, err := services.UserService.GetUserInfo(c.Keys["id"].(string)); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
