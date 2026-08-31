package handler

import (
	"go-pharma/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var medicines = []model.Medicine{
	{Id: 1, Name: "Paracetamol 500 mg", Category: "Analgesik", Stock: 120, Unit: "tablet", Price: 500},
	{Id: 2, Name: "Amoxicillin 500mg", Category: "Antibiotik", Stock: 45, Unit: "tablet", Price: 1500},
	{Id: 3, Name: "Betadine 60ml", Category: "Antiseptik", Stock: 30, Unit: "botol", Price: 12000},
}

func GetMedicine(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": medicines,
	})
}

func CreateMedicine(c *gin.Context) {
	var newMedicine model.Medicine
	if err := c.ShouldBindJSON(&newMedicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newMedicine.Id = len(medicines) + 1
	medicines = append(medicines, newMedicine)

	c.JSON(http.StatusCreated, gin.H{"data": newMedicine})
}
