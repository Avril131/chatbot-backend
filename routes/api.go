package routes

import (
	"chatbot-backend/app/controllers/app"
	"chatbot-backend/app/middleware"

	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes : deine api group routes
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/google", app.Login)

	authRouter :=router.Group("").Use(middleware.JWTAuth("web"))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
	}
}
