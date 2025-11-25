package controllers

import (
	"farhan_s/database"
	"farhan_s/entities"
	"farhan_s/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitProj(c *gin.Context) {
	err := repositories.UtkString(database.DBKonesi)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetStatus(c *gin.Context) {
	var result gin.H
	status, err := repositories.LihatStatus(database.DBKonesi)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": status,
		}
	}
	c.JSON(http.StatusOK, result)
}

func UpdateStatus(c *gin.Context) {
	var status entities.Status
	Ganti, _ := strconv.Atoi(c.Param("Ganti"))
	status.ServoStatus = Ganti
	err := repositories.UbahStatus(database.DBKonesi, status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ServoStatus": status.ServoStatus})
}
