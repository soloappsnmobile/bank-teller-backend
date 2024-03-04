package main

import (
	"bank-teller-backend/initializers"
	"bank-teller-backend/routers"
	"fmt"
	"log"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	router := routers.SetupRouter()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router.Run(fmt.Sprintf(":%s", port))
}
