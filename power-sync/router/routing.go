package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-sync/home"
	"power-sync/resources"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	{
		router.GET("/", home.Home)
		router.GET("/people/:id", resources.People)

		router.NoRoute(func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Route not found",
			})
		})
	}
	return router
}
