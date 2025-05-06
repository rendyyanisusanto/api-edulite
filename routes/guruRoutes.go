package routes

import (
	"edulite-api/controllers"
	"edulite-api/middleware"

	"github.com/gin-gonic/gin"
)

func GuruRoutes(r *gin.Engine) {
	guru := r.Group("/guru", middleware.AuthMiddleware())
	{
		guru.GET("/:id", controllers.GetGurubyId)
		guru.GET("/", controllers.GetAllGuru)
		// Tambah lagi POST, PUT, DELETE jika ada
	}
}
