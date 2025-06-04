package handler

import (
	"net/http"

	"github.com/EduardoMark/gym-api/internal/user"
	userPkg "github.com/EduardoMark/gym-api/internal/user"
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
	router.GET("/users/:id", h.FindOne)
	router.GET("/users", h.FindAll)
	router.PUT("/users/:id", h.Update)
	router.DELETE("/users/:id", h.Delete)
}

func (h *UserHandler) FindOne(c *gin.Context) {
	id := c.Param("id")

	user, err := h.useCase.FindOne(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": userPkg.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Phone:     user.Phone,
		Gender:    user.Gender,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}})
}

func (h *UserHandler) FindAll(c *gin.Context) {
	users, err := h.useCase.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := []userPkg.UserResponse{}
	for i := range users {
		response = append(response, userPkg.UserResponse{
			ID:        users[i].ID,
			Name:      users[i].Name,
			Email:     users[i].Email,
			Role:      users[i].Role,
			Phone:     users[i].Phone,
			Gender:    users[i].Gender,
			Address:   users[i].Address,
			CreatedAt: users[i].CreatedAt,
			UpdatedAt: users[i].UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"users": response})
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var body userPkg.UserRequest
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
