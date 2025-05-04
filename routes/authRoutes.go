package routes

import (
	"edulite-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/token", controllers.GenerateToken)
}
