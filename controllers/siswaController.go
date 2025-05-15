package controllers

import (
	"edulite-api/database"
	"edulite-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSiswaByNISN godoc
// @Summary Get siswa by NISN
// @Description Ambil data siswa berdasarkan NISN
// @Tags Siswa
// @Accept  json
// @Produce  json
// @Param nisn path string true "NISN"
// @Success 200 {object} models.Siswa
// @Failure 404 {object} map[string]string
// @Router /siswa/{nisn} [get]
// @Security BearerAuth
func GetSiswaByNISN(c *gin.Context) {
	db := database.DB
	nisn := c.Param("nisn")

	var siswa models.Siswa
	if err := db.Preload("Kelas").Preload("Jurusan").Where("nisn = ?", nisn).First(&siswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Siswa tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, siswa)
}

func GetAllSiswa(c *gin.Context) {
	db := database.DB

	var siswa []models.Siswa
	if err := db.Preload("Kelas").Preload("Jurusan").Limit(20).Find(&siswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, siswa)
}

func GetSiswaByID(c *gin.Context) {
	db := database.DB

	id := c.Param("id")

	var siswa models.Siswa
	if err := db.Preload("Kelas").Preload("Jurusan").First(&siswa, "id_siswa = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Siswa tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, siswa)
}
