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
func GetGurubyId(c *gin.Context) {
	db := database.DB
	id := c.Param("id")

	var guru models.Guru
	if err := db.Where("id = ?", id).First(&guru).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Guru tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, guru)
}

func GetAllGuru(c *gin.Context) {
	db := database.DB

	var guru []models.Guru
	if err := db.Where("is_active = 1").Limit(20).Find(&guru).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, guru)
}
