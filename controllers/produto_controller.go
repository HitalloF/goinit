package controllers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProdutoController struct {
	DB *gorm.DB
}

func NewProdutoController(db *gorm.DB) *ProdutoController {
	return &ProdutoController{
		DB: db,
	}
}

// Criar produto
func (ctrl *ProdutoController) CreateProduto(c *gin.Context) {
	var newProduto models.Produto

	if err := c.ShouldBindJSON(&newProduto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctrl.DB.Create(&newProduto)

	c.JSON(http.StatusCreated, newProduto)

}
