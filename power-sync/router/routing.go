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
		router.GET("/people", resources.GetStarWarsPeople)
		router.GET("/people/:name", resources.GetStarWarsPeopleById)
		router.POST("/people", resources.CreatePeople)
		router.PUT("/people/:name", resources.UpdatePeople)
		router.DELETE("/people/:name", resources.DeletePeople)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Route not found",
		})
	})
	return router
}
