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
	api.GET("/get", compHandler.GetUrl)

	adminRouter := api.Group("/admin")
	adminRouter.Use(middleware.AdminMiddleware())
	{
		adminRouter.POST("/register", compHandler.RegisterUrl)
	}

	likeRouter := api.Group("/like")
	{
		likeRouter.GET("/add", compHandler.AddLike)
	}
}
