package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiteshwagh9383/CodePulse/internal/model"
	"github.com/hiteshwagh9383/CodePulse/internal/repository"
	"github.com/hiteshwagh9383/CodePulse/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize the database connection
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	repository.InitDB()

	// Create the Gin router
	router := gin.Default()

	// Define routes
	router.GET("/tasks", getTasks)          // Get tasks from DB
	router.GET("/repos", fetchRepositories) // Get repositories from GitHub
	router.GET("/commits", fetchCommits)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Handler to fetch tasks from the database
func getTasks(c *gin.Context) {
	db := repository.GetDB()
	var tasks []model.Task
	if err := db.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func fetchCommits(c *gin.Context) {
	db := repository.GetDB()
	var repos []model.Repository
	db.Find(&repos)

	for _, repo := range repos {
		commits, err := service.FetchCommitsForRepo(repo.Name)
		if err != nil {
			log.Println("Failed to fetch commits:", err)
			continue
		}

		for _, commit := range commits {
			commit.RepositoryID = repo.ID
			commit.ID = commit.CommitHash

			result := db.FirstOrCreate(&commit, model.Commit{ID: commit.ID})
			if result.Error != nil {
				log.Println("Failed to insert commit:", err)
				continue
			}
			fmt.Println("Commit inserted or found successfully:", commit.CommitHash)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Commits fetched and stored"})
	}
}

// Handler to fetch repositories from GitHub
func fetchRepositories(c *gin.Context) {
	repos, err := service.FetchRepositories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch repositories"})
		return
	}
	db := repository.GetDB()
	var repos1 []model.Repository

	// Fetch all records
	if err := db.Find(&repos1).Error; err != nil {
		fmt.Println("Error fetching repositories:", err)
		return
	}
	for _, r := range repos {

		var existing model.Repository
		// FirstOrCreate looks for a record, and if not found, creates it
		result := db.FirstOrCreate(&existing, model.Repository{Name: r.Name, URL: r.URL})

		if result.Error != nil {
			fmt.Println("Failed to insert or find repository:", result.Error)
			continue
		}

		// After FirstOrCreate, 'existing' is either found or created.
		if result.RowsAffected == 0 {
			fmt.Println("No changes made (existing record found)", r.Name)
		} else {
			fmt.Println("Repository inserted or found successfully:", r.Name)
		}
	}

	c.JSON(http.StatusOK, repos)
}
