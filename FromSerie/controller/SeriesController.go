package controller

import (
	connector "FromSerie/connectors"
	"FromSerie/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllSeries(c *gin.Context) {
	var series []entities.Series

	connector.Db.Find(&series)
	c.JSON(http.StatusCreated, gin.H{
		"data": series,
	})
}

func CreateSeries(c *gin.Context) {
	var series entities.Series

	if err := c.ShouldBindJSON(&series); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	if err := connector.Db.Create(&series).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": series,
	})
}
