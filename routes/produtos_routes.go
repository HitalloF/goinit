package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProdutoRoutes(router *gin.Engine, produtoController *controllers.ProdutoController) {
	produtoRoutes := router.Group("/produtos")
	{
		produtoRoutes.POST("/", produtoController.CreateProduto)
	}
}
