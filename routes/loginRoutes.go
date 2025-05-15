package routes

import (
	"edulite-api/controllers"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(r *gin.Engine) {
	// Tanpa AuthMiddleware
	r.POST("/login", controllers.Login)
}
