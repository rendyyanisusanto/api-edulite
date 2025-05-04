package controllers

import (
	"edulite-api/database"
	"edulite-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSKLByCode godoc
// @Summary Ambil data SKL berdasarkan kode
// @Description Mengambil data alumni dan file SKL berdasarkan kode unik
// @Tags SKL
// @Accept  json
// @Produce  json
// @Param code path string true "Kode unik SKL"
// @Success 200 {object} models.Alumni
// @Failure 404 {object} map[string]string
// @Router /skl/{code} [get]
// @Security BearerAuth
func GetSKLByCode(c *gin.Context) {
	code := c.Param("code")
	db := database.DB

	var skl models.SKL
	err := db.Preload("Siswa").Where("kode = ? ", code).First(&skl).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data SKL tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, skl)
}
