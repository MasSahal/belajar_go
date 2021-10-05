package controllers

import (
	"mahasiswa/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InputMahasiswa struct {
	NIM      int    `json:"nim"`
	Nama     string `json:"nama"`
	Prodi    string `json:"prodi"`
	Semester int    `json:"semester"`
}

func AddMahasiswa(c *gin.Context) {

	//deklarasi databse
	db := c.MustGet("db").(*gorm.DB)

	var inputmahasiswa InputMahasiswa

	if err := c.ShouldBindJSON(&inputmahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ya": err.Error()})
		return
	}

	add := models.Mahasiswa{}
	add.NIM = inputmahasiswa.NIM
	add.Nama = inputmahasiswa.Nama
	add.Prodi = inputmahasiswa.Prodi
	add.Semester = inputmahasiswa.Semester

	err := db.Create(&add).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed create mahasiswa!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": add})
}

func UpdateMahasiswa(c *gin.Context) {

	//deklarasi databse
	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("id")
	var mahasiswa models.Mahasiswa

	//cari data mahasiswa - perhatikan .Error
	if err := db.Where("id = ?", id).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  404,
			"error":   err.Error(),
			"message": "Data mahasiswa tidak ditemukan!",
		})
		return
	}

	var inputmahasiswa InputMahasiswa
	//validasi input
	if err := c.ShouldBindJSON(&inputmahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Data mahasiswa gagal di eksekusi JSON!",
		})
		return
	}

	up := models.Mahasiswa{}
	up.NIM = inputmahasiswa.NIM
	up.Nama = inputmahasiswa.Nama
	up.Prodi = inputmahasiswa.Prodi
	up.Semester = inputmahasiswa.Semester

	err := db.Model(&mahasiswa).Updates(&up).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed create mahasiswa!"})
	}

	data := db.Where("id = ?", id).First(&mahasiswa).Error
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  200,
		"message": "Data mahasiswa ditemukan!",
	})
}

func DeleteMahasiswa(c *gin.Context) {

	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var mahasiswa []models.Mahasiswa

	//cari data
	err := db.Where("id = ?", id).Delete(&mahasiswa).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}
	//berikan hasil
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Data mahasiswa berhasil dihapus!",
	})
}

func GetMahasiswa(c *gin.Context) {

	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var mahasiswa []models.Mahasiswa

	//cari data
	err := db.Where("id = ?", id).First(&mahasiswa).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
		return
	}

	//berikan feedback
	c.JSON(http.StatusOK, gin.H{
		"data":    mahasiswa,
		"status":  "200",
		"message": "Data mahasiswa berhasil ditemukan!",
	})
}

func GetAllMahasiswa(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var mahasiswa []models.Mahasiswa

	//cari data
	db.Find(&mahasiswa)

	//berikan hasil
	c.JSON(http.StatusOK, gin.H{
		"data":    mahasiswa,
		"status":  "200",
		"message": "Data berhasil ditemukan!",
	})
}

// func UpdateMahasiswa(c *gin.Context) {

// 	id := c.Param("id")
// 	db := c.MustGet("db").(*gorm.DB)
// 	var mahasiswa []models.Mahasiswa

// 	//cari data
// 	err := db.Where("id = ?", id).First(&mahasiswa).Error

// 	//ambil inputan

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!"})
// 		return
// 	}
// 	//berikan hasil
// 	c.JSON(http.StatusOK, gin.H{
// 		"data":    mahasiswa,
// 		"status":  "200",
// 		"message": "Data mahasiswa berhasil ditemukan!",
// 	})
// }
