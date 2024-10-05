package main

import (
	"todogorest/config"
	"todogorest/routes"
)

func main() {
	config.InitDB()           // Initialize database connection
	r := routes.SetupRouter() // Setup routes
	r.Run(":8080")            // Run the application on port 8080
}
