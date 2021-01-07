package routes

import (
	"cubicasa/configs"
	"cubicasa/controllers"
	"cubicasa/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(r *gin.Engine) {
	// swagger info
	docs.SwaggerInfo.Title = "Web Service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("127.0.0.1:%v", configs.Port)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	userController := controllers.UserController{}
	r.POST("/user", userController.Create)
	r.POST("/user/assign-to-team", userController.AssignTeam)

	teamController := controllers.TeamController{}
	r.POST("/team", teamController.Create)
	r.POST("/team/assign-to-hub", teamController.AssignHub)

	hubController := controllers.HubController{}
	r.POST("/hub", hubController.Create)

	searchController := controllers.SearchController{}
	r.GET("/search", searchController.Search)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
