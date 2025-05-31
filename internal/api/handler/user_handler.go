package handler

import (
	"net/http"

	"github.com/EduardoMark/gym-api/internal/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	useCase user.UseCase
}

func NewUserHandler(uc user.UseCase) *UserHandler {
	return &UserHandler{
		useCase: uc,
	}
}

func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/users", h.Create)
	router.GET("/users/:id", h.FindOne)
	router.GET("/users", h.FindAll)
	router.PUT("/users/:id", h.Update)
	router.DELETE("/users/:id", h.Delete)
}

func (h *UserHandler) Create(c *gin.Context) {
	var body user.UserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.useCase.Create(body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *UserHandler) FindOne(c *gin.Context) {
	id := c.Param("id")

	record, err := h.useCase.FindOne(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user.UserResponse{
		ID:        record.ID,
		Name:      record.Name,
		Email:     record.Email,
		Role:      record.Role,
		Phone:     record.Phone,
		Gender:    record.Gender,
		Address:   record.Address,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}})
}

func (h *UserHandler) FindAll(c *gin.Context) {
	records, err := h.useCase.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := make([]user.UserResponse, len(records))
	for i, record := range records {
		response[i] = user.UserResponse{
			ID:        record.ID,
			Name:      record.Name,
			Email:     record.Email,
			Role:      record.Role,
			Phone:     record.Phone,
			Gender:    record.Gender,
			Address:   record.Address,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{"users": response})
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var body user.UserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	if err := h.useCase.Update(id, body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Success on update user"})
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.useCase.Delete(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "user deleted with success"})
}
