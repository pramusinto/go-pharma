package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"go-pharma/internal/model"
	"go-pharma/internal/repository"

	"github.com/gin-gonic/gin"
)

type MedicineHandler struct {
	repo *repository.MedicineRepository
}

func NewMedicineHandler(repo *repository.MedicineRepository) *MedicineHandler {
	return &MedicineHandler{repo: repo}
}

func (h *MedicineHandler) GetMedicines(c *gin.Context) {
	search := c.Query("search") // ?search=paracetamol

	medicines, err := h.repo.GetAll(context.Background(), search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicines})
}

func (h *MedicineHandler) CreateMedicine(c *gin.Context) {
	var m model.Medicine
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Create(context.Background(), &m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": m})
}

func (h *MedicineHandler) UpdateMedicine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var m model.Medicine
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Update(context.Background(), id, &m); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "medicine not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	m.Id = id
	c.JSON(http.StatusOK, gin.H{"data": m})
}

func (h *MedicineHandler) DeleteMedicine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.repo.Delete(context.Background(), id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "medicine not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
