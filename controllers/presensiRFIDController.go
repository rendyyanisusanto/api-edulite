package controllers

import (
	"edulite-api/database"
	"edulite-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPresenceToday(c *gin.Context) {
	db := database.DB
	idsiswa_fk := c.Param("id")

	var presensi_rfid []models.PresensiRFID
	if err := db.Preload("Siswa").Where("idsiswa_fk = ?", idsiswa_fk).Find(&presensi_rfid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Gagal mengambil data"})
		return
	}

	if len(presensi_rfid) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []models.PresensiRFID{}, "message": "Tidak ada data presensi hari ini"})
		return
	}

	c.JSON(http.StatusOK, presensi_rfid)
}

func GetPresenceInToday(c *gin.Context) {
	db := database.DB
	idsiswa_fk := c.Param("id")

	var presensi_rfid []models.PresensiRFID
	if err := db.Where("idsiswa_fk = ?", idsiswa_fk).Where("status = 'MASUK'").First(&presensi_rfid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Gagal mengambil data"})
		return
	}

	if len(presensi_rfid) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []models.PresensiRFID{}, "message": "Tidak ada data presensi hari ini"})
		return
	}

	c.JSON(http.StatusOK, presensi_rfid)
}

func GetPresenceOutToday(c *gin.Context) {
	db := database.DB
	idsiswa_fk := c.Param("id")

	var presensi_rfid []models.PresensiRFID
	if err := db.Where("idsiswa_fk = ?", idsiswa_fk).Where("status = 'PULANG'").Last(&presensi_rfid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Gagal mengambil data"})
		return
	}

	if len(presensi_rfid) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []models.PresensiRFID{}, "message": "Tidak ada data presensi hari ini"})
		return
	}

	c.JSON(http.StatusOK, presensi_rfid)
}
