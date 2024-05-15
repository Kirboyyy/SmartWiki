package main

import (
	"github.com/kirboyyy/smartwiki/internal/adapters/repositories/sqlite"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kirboyyy/smartwiki/config"
	"github.com/kirboyyy/smartwiki/database"
	"github.com/kirboyyy/smartwiki/internal/adapters/controllers"
	"github.com/kirboyyy/smartwiki/internal/service"
)

func main() {
	config.LoadConfig()
	database.InitDB()
	pageRepo := &sqlite.PageRepository{DB: database.DB}
	pageService := &service.PageService{PageRepo: pageRepo}
	pageController := &controllers.PageController{PageService: pageService}

	router := gin.Default()
	router.GET("/api/page/:id", pageController.GetPageHandler)
	router.POST("/api/page", pageController.CreatePageHandler)
	router.PUT("/api/pages/:id", pageController.UpdatePageHandler)
	router.DELETE("/api/pages/:id", pageController.DeletePageHandler)

	log.Println("Server starting on :8080")
	router.Run(":8080")
}
