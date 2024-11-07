package resources

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func People(c *gin.Context) {

	c.Query("")
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"people": []string{"Luke skywalker", "Anakin skywalker", "Darth vader"},
	})
}
