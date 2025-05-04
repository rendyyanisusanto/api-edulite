package routes

import (
	"edulite-api/controllers"
	"edulite-api/middleware"

	"github.com/gin-gonic/gin"
)

func SklRoutes(r *gin.Engine) {
	protected := r.Group("/skl", middleware.AuthMiddleware())
	{
		protected.GET("/:code", controllers.GetSKLByCode)
	}
}
