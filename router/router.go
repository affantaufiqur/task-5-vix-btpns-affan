package router

import (
	"affan/controllers"
	"affan/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	
	r.POST("/user/register", controllers.UserRegister)
	r.POST("/user/login", controllers.UserLogin)
	r.PUT("/user/:userId", controllers.UserUpdate)
	r.DELETE("/user/:userId", controllers.UserDelete)

	r.Use(middlewares.Authentication())
	r.POST("/photo/upload", controllers.CreatePhoto)
	r.GET("/photo", controllers.ListPhoto)
	r.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
	r.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	
	return r
}