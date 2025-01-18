package services

import (
	"chatbot-backend/app/common/request"
	"chatbot-backend/app/models"
	"chatbot-backend/global"
	"context"
	"errors"
	"fmt"
	"strconv"

	"google.golang.org/api/idtoken"
)

type userService struct {
}

var UserService = new(userService)

func (userService *userService) Login(params request.Login) (user *models.User, err error) {
	clientID := global.App.Config.GoogleID

	// validate user token
	payload, err := idtoken.Validate(context.Background(), params.Token, clientID)
	if err != nil {
		return
	}

	// find user info
	var result = global.App.DB.Where("g_id = ?", payload.Claims["aud"].(string)).First(&user)
	if result.RowsAffected == 0{
		user = &models.User{
			GId:     payload.Claims["aud"].(string),
			Email:   payload.Claims["email"].(string),
			Picture: payload.Claims["picture"].(string),
			Name:    payload.Claims["name"].(string),
		}

		// insert new user into db
		err := global.App.DB.Create(&user).Error
		if err != nil {
			return user, fmt.Errorf("failed to create user: %v", err)
		}
	}

	return user, nil
}

func (userService *userService) GetUserInfo(id string) (user models.User, err error) {
	intID, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intID).Error
	if err != nil {
		err = errors.New("user not exists")
	}
	return
}
