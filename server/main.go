package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiteshwagh9383/CodePulse/internal/model"
	"github.com/hiteshwagh9383/CodePulse/internal/repository"
)

func main() {
	// Initialize the database connection
	repository.InitDB()

	// Create the Gin router
	router := gin.Default()

	// Define routes
	router.GET("/tasks", getTasks)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getTasks(c *gin.Context) {
	db := repository.GetDB()

	var tasks []model.Task
	if err := db.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
