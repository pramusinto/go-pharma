package main

import (
	"go-pharma/internal/handler"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	api := r.Group("/api")
	{
		api.GET("/medicines", handler.GetMedicine)
		api.POST("/medicines", handler.CreateMedicine)
	}

	log.Println("Server Jalan di http://localhost:8080")
	r.Run(":8080")
}
