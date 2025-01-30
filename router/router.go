package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shariquehaider/ecom-backend/controllers"
	"github.com/shariquehaider/ecom-backend/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/api/register", controllers.RegisterController)
	router.POST("/api/login", controllers.LoginController)
	router.GET("/api/account", middleware.VerifyTokenMiddleware(), controllers.GetProfileController)
	return router
}
