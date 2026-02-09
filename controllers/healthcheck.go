package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckRoutes(rg *gin.RouterGroup) {
	health := rg.Group("/health")
	{
		health.GET("/", HealthChecking)
	}
}

func HealthChecking(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Golang Gin Application is running successfully",
	})
}
