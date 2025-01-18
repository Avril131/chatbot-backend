package routes

import (
	"chatbot-backend/app/controllers/app"
	"chatbot-backend/app/middleware"

	"github.com/gin-gonic/gin"
)

// SetAPIGroupRoutes : deine api group routes
func SetAPIGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/google", app.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth("web"))
	{
		// get user basic info
		authRouter.POST("/auth/info", app.Info)

		// get chat list by user
		authRouter.POST("/chat-list", app.GetChatList)

		// get history message by chat id
		authRouter.POST("/messages", app.GetMessages)

		// create a new chat
		authRouter.POST("/create-chat", app.CreateChat)

		// send message to get res
		authRouter.POST("/send", app.SendMsg)

		// log out
		authRouter.POST("/auth/logout", app.Logout)
	}
}
