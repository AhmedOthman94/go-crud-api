package main

import (
	"log"

	"github.com/AhmedOthman94/go-crud-api/initializers"
	"github.com/AhmedOthman94/go-crud-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	routes.PostRoutes(r)
	r.Run()
}
