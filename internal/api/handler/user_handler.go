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
}

func (h *UserHandler) Create(c *gin.Context) {
	var user user.CreateUserParams
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.useCase.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
