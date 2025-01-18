package services

import (
	"chatbot-backend/app/common/request"
	"chatbot-backend/app/models"
	"chatbot-backend/global"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type chatService struct {
}

var ChatServices = new(chatService)

func (ChatService *chatService) GetChatListByUID(id string) (chats []models.Chat, err error) {
	intID, err := strconv.Atoi(id)

	err = global.App.DB.Where("u_id = ?", intID).Order("updated_at DESC").Find(&chats).Error
	if err != nil {
		err = errors.New("something wrong in finding data")
	}

	return
}

func (ChatService *chatService) CheckChatByUID(id string, params request.QueryMessage) (b bool, err error) {

	var chat models.Chat
	result := global.App.DB.Where("id = ? AND u_id = ?", params.CID, id).First(&chat)

	// check the result of search
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {

			b = false
		}
		err = result.Error
	} else {
		b = true
	}

	return
}

func (ChatService *chatService) GetMessagesByCID(params request.QueryMessage) (messages []models.Message, err error) {

	result := global.App.DB.Where("c_id = ?", params.CID).
		Order("created_at DESC").
		Find(&messages) // save result into mesaages

	if result.Error != nil {
		return nil, result.Error
	}

	return
}

func (ChatService *chatService) CreateChat(id string) (chat models.Chat, err error) {

	uid, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	chat = models.Chat{
		Name: "New Chat",
		UserID:  uid,
	}

	// insert into db
	result := global.App.DB.Create(&chat)

	if result.Error != nil {
		return chat, result.Error
	}

	return chat, nil
}
