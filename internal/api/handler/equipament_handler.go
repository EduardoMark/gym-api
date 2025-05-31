package handler

import (
	"net/http"

	"github.com/EduardoMark/gym-api/internal/equipament"
	"github.com/gin-gonic/gin"
)

type EquipamentHandler struct {
	uc equipament.UseCase
}

func NewEquipamentHandler(uc equipament.UseCase) *EquipamentHandler {
	return &EquipamentHandler{uc: uc}
}

func (h EquipamentHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/equipament", h.Create)
	router.GET("/equipament/:id", h.FindOne)
	router.GET("/equipament", h.FindAll)
	router.PUT("/equipament/:id", h.Update)
	router.DELETE("/equipament/:id", h.Delete)
}

func (h *EquipamentHandler) Create(c *gin.Context) {
	var body equipament.EquipamentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Create(body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "equipament created with success"})
}

func (h *EquipamentHandler) FindOne(c *gin.Context) {
	id := c.Param("id")

	record, err := h.uc.FindOne(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := equipament.EquipamentResponse{
		ID:              record.ID,
		Name:            record.Name,
		Description:     record.Description,
		Category:        record.Category,
		Brand:           record.Brand,
		Model:           record.Model,
		MaintenanceDate: record.MaintenanceDate,
		Status:          record.Status,
		Quantity:        record.Quantity,
		CreatedAt:       record.CreatedAt,
		UpdatedAt:       record.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"equipament": response})
}

func (h *EquipamentHandler) FindAll(c *gin.Context) {
	records, err := h.uc.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := make([]equipament.EquipamentResponse, len(records))
	for i, record := range records {
		response[i] = equipament.EquipamentResponse{
			ID:              record.ID,
			Name:            record.Name,
			Description:     record.Description,
			Category:        record.Category,
			Brand:           record.Brand,
			Model:           record.Model,
			MaintenanceDate: record.MaintenanceDate,
			Status:          record.Status,
			Quantity:        record.Quantity,
			CreatedAt:       record.CreatedAt,
			UpdatedAt:       record.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{"equipaments": response})
}

func (h *EquipamentHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var body equipament.EquipamentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Update(id, body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "user updated with success"})
}

func (h *EquipamentHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.uc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "equipament deleted with success"})
}
