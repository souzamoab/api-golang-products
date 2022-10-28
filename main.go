package main

import (
	"github.com/fabriciojlm/api-golang-users/src/app/database"
	_ "github.com/fabriciojlm/api-golang-users/src/app/docs"
	"github.com/fabriciojlm/api-golang-users/src/app/server/routes"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @host      localhost:8090
// @BasePath  /api/v1/users
func main() {
	database.StartDatabase()
	godotenv.Load()
	router := routes.ConfigRoutes()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
