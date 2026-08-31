package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pramusinto/go-pharma/backend/internal/config"
	"github.com/pramusinto/go-pharma/backend/internal/handler"
	"github.com/pramusinto/go-pharma/backend/internal/repository"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	medicineRepo := repository.NewMedicineRepository(db)
	medicineHandler := handler.NewMedicineHandler(medicineRepo)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	api := r.Group("/api")
	{
		api.GET("/medicines", medicineHandler.GetMedicines)
		api.POST("/medicines", medicineHandler.CreateMedicine)
		api.PUT("/medicines/:id", medicineHandler.UpdateMedicine)
		api.DELETE("/medicines/:id", medicineHandler.DeleteMedicine)
	}

	log.Println("Server jalan di http://localhost:8080")
	r.Run(":8080")
}
