package routers

import (
	"FromSerie/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainRouter(router *gin.Engine) *gin.Engine {
	api := router.Group("/v1/series")
	{
		api.GET("/health-check", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Server is up and running",
			})
		})
		api.GET("/", controller.GetAllSeries)
		api.POST("/", controller.CreateSeries)
		api.GET("/characters", controller.CharacterController)
	}
	return router
}
