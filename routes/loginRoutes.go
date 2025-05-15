package routes

import (
	"edulite-api/controllers"
	"edulite-api/middleware"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(r *gin.Engine) {
	protected := r.Group("/login")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/", controllers.Login)
	}
}
