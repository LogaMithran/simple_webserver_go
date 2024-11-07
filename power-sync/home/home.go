package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "Hello world!",
		"message": "/ controller works",
	})
}
