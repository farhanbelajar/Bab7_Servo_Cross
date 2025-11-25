package routers

import (
	"farhan_s/controllers"
	"time"

	"github.com/gin-contrib/cors" // sama ini
	"github.com/gin-gonic/gin"
)

func MulaiServer() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/servo/init-proj", controllers.InitProj)
	router.GET("/servo/status", controllers.GetStatus)
	router.PUT("/servo/update/:Ganti", controllers.UpdateStatus)
	return router
}
