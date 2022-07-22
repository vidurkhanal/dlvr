package main

import (
	"auth-and-users/controllers"
	"auth-and-users/repository"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.Connect(os.Getenv("POSTGRES_URI"))
	repository.Migrate()
	r := initRouter()
	// _httpPort := os.Getenv("HTTP_PORT")

	r.Run(":8060")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		api.GET("/user/who-am-i", controllers.WhoAmI)
		// secured := api.Group("/secured").Use(middlewares.Auth())
		// {
		// 	secured.GET("/ping", controllers.Ping)
		// }
	}
	return router
}
