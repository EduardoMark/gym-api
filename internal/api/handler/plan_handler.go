package handler

import (
	"net/http"

	"github.com/EduardoMark/gym-api/internal/payment/plan"
	"github.com/gin-gonic/gin"
)

type PlanHandler struct {
	uc plan.UseCase
}

func NewPlanHandler(uc plan.UseCase) *PlanHandler {
	return &PlanHandler{uc: uc}
}

func (h *PlanHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/plans", h.Create)
	router.GET("/plans/:id", h.FindOne)
	router.GET("/plans", h.FindAll)
	router.PUT("/plans/:id", h.Update)
	router.DELETE("/plans/:id", h.Delete)
}

func (h *PlanHandler) Create(c *gin.Context) {
	var body plan.PlanRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Create(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Plan created with success"})
}

func (h *PlanHandler) FindOne(c *gin.Context) {
	id := c.Param("id")

	record, err := h.uc.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := plan.PlanResponse{
		ID:               record.ID,
		Name:             record.Name,
		Description:      record.Description,
		Price:            record.Price,
		Cicle:            record.Cicle,
		DurationInMonths: record.DurationInMonths,
		CreatedAt:        record.CreatedAt,
		UpdatedAt:        record.UpdatedAt,
	}

	c.JSON(http.StatusOK, res)
}

func (h *PlanHandler) FindAll(c *gin.Context) {
	records, err := h.uc.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := make([]plan.PlanResponse, len(records))

	for i, record := range records {
		res[i] = plan.PlanResponse{
			ID:               record.ID,
			Name:             record.Name,
			Description:      record.Description,
			Price:            record.Price,
			Cicle:            record.Cicle,
			DurationInMonths: record.DurationInMonths,
			CreatedAt:        record.CreatedAt,
			UpdatedAt:        record.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, res)
}

func (h *PlanHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var body plan.PlanRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Update(id, &body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Plan updated with success"})
}

func (h *PlanHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.uc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Plan deleted with success"})
}
