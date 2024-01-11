package controllers

import (
	"main/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		DB: db,
	}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctrl.DB.Create(&newUser)

	c.JSON(http.StatusCreated, newUser)
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	result := ctrl.DB.First(&user, userID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) ListUsers(c *gin.Context) {
	var users []models.User
	ctrl.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}

// Deletar User
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var existingUser models.User
	result := ctrl.DB.First(&existingUser, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	ctrl.DB.Delete(&existingUser)

	c.JSON(http.StatusNoContent, nil)
}

// Atualizar user
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var existingUser models.User
	result := ctrl.DB.First(&existingUser, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&existingUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctrl.DB.Save(&existingUser)

	c.JSON(http.StatusOK, existingUser)
}
