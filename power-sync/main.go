package main

import (
	"fmt"
	"os"
	"power-sync/connectors"
	"power-sync/router"
)

func main() {
	connectors.CreateDataBaseConnection()
	r := router.SetupRouter()

	if err := r.Run("localhost:8080"); err != nil {
		fmt.Println("Error in starting the server")
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		err := connectors.Close()
		if err != nil {
			fmt.Println("Error in closing the database")
		}
	}()
}
