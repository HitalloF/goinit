package main

import (
	"fmt"
	"main/controllers"
	"main/database"
	"main/models"
	"main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	database.Init()
	models.MigrateUser(database.Init())
	router := gin.Default()
	userController := controllers.NewUserController(database.Init())
	produtorController := controllers.NewProdutoController(database.Init())
	routes.SetupUserRoutes(router, userController)
	routes.SetupProdutoRoutes(router, produtorController)
	router.Run(":8180")
}
