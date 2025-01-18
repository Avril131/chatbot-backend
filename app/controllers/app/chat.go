package app

import (
	"chatbot-backend/app/common/request"
	"chatbot-backend/app/common/response"
	"chatbot-backend/app/services"

	"github.com/gin-gonic/gin"
)

// GetChatList get chat list
func GetChatList(c *gin.Context) {
	if chats, err := services.ChatServices.GetChatListByUID(c.Keys["id"].(string)); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, chats)
	}
}

// GetMessages get message by chat id
func GetMessages(c *gin.Context) {

	uid := c.Keys["id"].(string)

	var form request.QueryMessage
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if isMatch, err := services.ChatServices.CheckChatByUID(uid, form); err != nil {
		response.BusinessFail(c, err.Error())
	} else if isMatch == false {
		response.TokenFail(c)
	}

	if messages, err := services.ChatServices.GetMessagesByCID(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, messages)
	}
}

// CreateChat create a new chat record
func CreateChat(c *gin.Context) {

	uid := c.Keys["id"].(string)

	if chat, err := services.ChatServices.CreateChat(uid); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, chat)
	}
}

// SendMsg handle message
func SendMsg(c *gin.Context) {

	var form request.Message
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	res, err := services.OpenAI.SendMessage(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, res)
	}
}
