package bootstrap

import (
	"chatbot-backend/app/middleware"
	"chatbot-backend/global"
	"chatbot-backend/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	if global.App.Config.App.Env == "production" {
        gin.SetMode(gin.ReleaseMode)
    }
    router := gin.New()
    router.Use(gin.Logger(), middleware.CustomRecovery())

	router.Use(middleware.Cors())

	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

// start server
func RunServer() {
	r := setupRouter()
	r.Run(":" + global.App.Config.App.Port)
}
