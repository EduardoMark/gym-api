package handler

import (
	"net/http"

	"github.com/EduardoMark/gym-api/internal/user"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc user.UseCase
}

func NewAuthHandler(uc user.UseCase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h *AuthHandler) RegisterRouter(router *gin.RouterGroup) {
	router.POST("/login", h.Login)
	router.POST("/signup", h.Signup)
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var user user.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.uc.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var body user.UserLoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.uc.Login(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
