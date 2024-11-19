package main

import (
	connector "FromSerie/connectors"
	"FromSerie/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	connector.InitializeDBConnection()
	router = routers.MainRouter(router)

	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "No routes matched",
		})
	})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := connector.CloseDbConnection()
		if err != nil {
			fmt.Println("Error in closing the database")
		}
	}()
}
