package main

import (
	"github.com/joho/godotenv"
	"github.com/souzamoab/api-golang-products/src/app/database"
	"github.com/souzamoab/api-golang-products/src/app/server/routes"
)

func main() {
	database.StartDatabase()
	godotenv.Load()
	router := routes.ConfigRoutes()
	router.Run(":8080")
}
