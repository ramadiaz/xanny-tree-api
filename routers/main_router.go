package routers

import (
	"xanny-tree-api/config"
	"xanny-tree-api/handlers"
	"xanny-tree-api/middleware"
	"xanny-tree-api/repositories"
	"xanny-tree-api/services"

	"github.com/gin-gonic/gin"
)

func CompRouter(api *gin.RouterGroup) {
	api.Use(middleware.ClientTracker(config.InitDB()))

	compRepository := repositories.NewComponentRepository(config.InitDB())
	compService := services.NewService(compRepository)
	compHandler := handlers.NewCompHandlers(compService)

	api.GET("/ping", compHandler.Ping)
	api.POST("/register", compHandler.RegisterUrl)
}
