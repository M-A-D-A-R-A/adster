package main

import (
	"log"
"microservices/core/database"
"microservices/core/routes"
"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseInstance := database.NewSupabase()

	routes.NewRoutes(supabaseInstance).RunAppRouter()

}
