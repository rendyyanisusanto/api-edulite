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
	if err := db.Where("idsiswa_fk = ?", idsiswa_fk).Where("status = 'MASUK'").Last(&presensi_rfid).Error; err != nil {
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

func GetPresenceMonthlyByStudent(c *gin.Context) {
	db := database.DB
	idsiswa_fk := c.Param("id")

	var presensi_rfid []models.PresensiRFID
	if err := db.Where("idsiswa_fk = ?", idsiswa_fk).Find(&presensi_rfid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Gagal mengambil data"})
		return
	}

	if len(presensi_rfid) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []models.PresensiRFID{}, "message": "Tidak ada data presensi hari ini"})
		return
	}

	c.JSON(http.StatusOK, presensi_rfid)
}
func GetPresenceMonthlyByStudentFormated(c *gin.Context) {
	db := database.DB
	idsiswa_fk := c.Param("id")

	var rekap []models.RekapPresensi
	query := `
	SELECT
		tanggal,
		idsiswa_fk,
		MIN(CASE WHEN status = 'MASUK' THEN waktu END) AS masuk,
		MAX(CASE WHEN status = 'PULANG' THEN waktu END) AS pulang,
		COUNT(CASE WHEN status = 'IJIN KELUAR' THEN 1 END) AS jumlah_ijin_keluar
	FROM
		presensi_rfid
	WHERE
		idsiswa_fk = ?
	GROUP BY
		tanggal, idsiswa_fk
	ORDER BY
		tanggal ASC
	`

	if err := db.Raw(query, idsiswa_fk).Scan(&rekap).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data rekap", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rekap})
}

func GetPresenceMonthlyCountByStudent(c *gin.Context) {
	db := database.DB
	idsiswa_fk := c.Param("id")

	var rekap []models.JumlahPresensi
	query := `
	SELECT COUNT(DISTINCT tanggal) as jumlah_masuk
	FROM presensi_rfid
	WHERE
	idsiswa_fk = ?
	AND MONTH(tanggal) = MONTH(CURDATE())
	AND YEAR(tanggal) = YEAR(CURDATE());
	`

	if err := db.Raw(query, idsiswa_fk).First(&rekap).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengambil data rekap", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rekap})
}
