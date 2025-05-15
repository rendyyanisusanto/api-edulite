package routes

import (
	"edulite-api/controllers"
	"edulite-api/middleware"

	"github.com/gin-gonic/gin"
)

func SiswaRoutes(r *gin.Engine) {
	siswa := r.Group("/siswa", middleware.AuthMiddleware())
	{
		siswa.GET("/id/:id", controllers.GetSiswaByID)
		siswa.GET("/nisn/:nisn", controllers.GetSiswaByNISN)
		siswa.GET("/", controllers.GetAllSiswa)

		// Tambah lagi POST, PUT, DELETE jika ada
	}
}
