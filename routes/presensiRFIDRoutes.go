package routes

import (
	"edulite-api/controllers"
	"edulite-api/middleware"

	"github.com/gin-gonic/gin"
)

func PresensiRFIDRoutes(r *gin.Engine) {
	presensiRFID := r.Group("/presensiRFID", middleware.AuthMiddleware())
	{
		presensiRFID.GET("/id/:id", controllers.GetPresenceToday)
		presensiRFID.GET("/IN/:id", controllers.GetPresenceInToday)
		presensiRFID.GET("/OUT/:id", controllers.GetPresenceOutToday)
		presensiRFID.GET("/PresenceMonthlyStudent/:id", controllers.GetPresenceMonthlyByStudent)
		presensiRFID.GET("/PresenceMonthlyStudentFormated/:id", controllers.GetPresenceMonthlyByStudentFormated)
		presensiRFID.GET("/PresenceMonthlyCountByStudent/:id", controllers.GetPresenceMonthlyCountByStudent)
		// Tambah lagi POST, PUT, DELETE jika ada
	}
}
