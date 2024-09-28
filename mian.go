package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yashbalyan08/shipment-app/config"
	"github.com/yashbalyan08/shipment-app/routes"
)

func main() {
	// Initialize the database connection
	config.InitDB()

	// Set up the routes
	router := routes.SetupRoutes(config.DB)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
