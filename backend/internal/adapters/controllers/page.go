package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kirboyyy/smartwiki/internal/domain"
	"github.com/kirboyyy/smartwiki/internal/service"
)

type PageController struct {
	PageService *service.PageService
}

func (pc *PageController) GetPageHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	page, err := pc.PageService.GetPage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}
	c.JSON(http.StatusOK, page)
}

func (pc *PageController) CreatePageHandler(c *gin.Context) {
	var page domain.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := pc.PageService.CreatePage(&page); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create page"})
		return
	}
	c.JSON(http.StatusOK, page)
}

func (pc *PageController) UpdatePageHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}

	var page domain.Page
	if err := c.ShouldBind(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	page.ID = id
	if _, err := pc.PageService.UpdatePage(&page); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update page"})
		return
	}
	c.JSON(http.StatusOK, page)
}

func (pc *PageController) DeletePageHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page ID"})
		return
	}
	err = pc.PageService.DeletePageById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete page"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
