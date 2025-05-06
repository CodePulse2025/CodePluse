package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiteshwagh9383/CodePulse/internal/model"
	"github.com/hiteshwagh9383/CodePulse/internal/repository"
	"github.com/hiteshwagh9383/CodePulse/internal/service"
	metrics "github.com/hiteshwagh9383/CodePulse/internal/service/metrics"
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
	// 👇 Add CORS middleware here
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // allow all origins
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
	// Define routes
	router.GET("/repos", fetchRepositories) // Get repositories from GitHub
	router.GET("/commits", fetchCommits)
	router.GET("/prs", fetchPRs) // Get pull requests from GitHub
	router.GET("/metrics/top-committers", GetTopCommitters)
	router.GET("/metrics/top-contributors", GetTopContributorsHandler)
	router.GET("/metrics/top-reviewers", GetTopReviewers)
	router.GET("/metrics/prs-open-longer-than-5-days", GetPRsOpenLongerThan5Days)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

type TopCommitter struct {
	Name  string
	Email string
	Count int
}

// Dummy handler for PRs open longer than 5 days
func GetPRsOpenLongerThan5Days(c *gin.Context) {
	// Dummy data representing PRs that are open for more than 5 days
	dummyData := []map[string]interface{}{
		{
			"repo_name": "Repo1",
			"pr_number": 123,
			"opened_at": "2025-04-25",
			"days_open": 7,
			"developer": "Alice",
			"pr_title":  "Feature X",
		},
		{
			"repo_name": "Repo2",
			"pr_number": 456,
			"opened_at": "2025-04-20",
			"days_open": 10,
			"developer": "Bob",
			"pr_title":  "Bugfix Y",
		},
		{
			"repo_name": "Repo3",
			"pr_number": 789,
			"opened_at": "2025-04-18",
			"days_open": 12,
			"developer": "Charlie",
			"pr_title":  "Improvement Z",
		},
	}

	// Return the dummy data
	c.JSON(http.StatusOK, dummyData)
}
func GetTopReviewers(c *gin.Context) {
	reviewers, err := metrics.GetTopReviewers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch top reviewers"})
		return
	}

	c.JSON(http.StatusOK, reviewers)
}
func GetTopContributorsHandler(c *gin.Context) {
	contributors, err := metrics.GetTopContributors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch top contributors"})
		return
	}

	c.JSON(http.StatusOK, contributors)
}

func GetTopCommitters(c *gin.Context) {
	db := repository.GetDB()
	var results []TopCommitter

	err := db.Table("commits").
		Select("developers.name, developers.email, COUNT(commits.id) as count").
		Joins("JOIN developers ON developers.id = commits.committer_id").
		Group("developers.name, developers.email").
		Order("count DESC").
		Limit(3).
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch top committers"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func fetchPRs(c *gin.Context) {
	db := repository.GetDB()
	var repos []model.Repository
	if err := db.Find(&repos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch repositories"})
		return
	}

	for _, repo := range repos {
		if err := service.FetchPullRequestsForRepo(repo.Name); err != nil {
			log.Println("Failed to fetch PRs:", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "PRs fetched and stored"})
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
