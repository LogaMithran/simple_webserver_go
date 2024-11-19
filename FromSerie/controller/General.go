package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GeneralInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message":   "Hello there!, this a series api of a tv show called from",
		"character": "/character",
	})
}
