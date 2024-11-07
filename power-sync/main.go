package main

import (
	"fmt"
	"os"
	"power-sync/router"
)

func main() {
	r := router.SetupRouter()

	if err := r.Run("localhost:8080"); err != nil {
		fmt.Println("Error in starting the server")
		fmt.Println(err)
		os.Exit(1)
	}
}
