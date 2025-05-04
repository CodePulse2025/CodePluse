package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hiteshwagh9383/CodePulse/internal/model"
	"github.com/hiteshwagh9383/CodePulse/internal/repository"
)

func FetchPullRequestsForRepo(repoName string) error {
	db := repository.GetDB()

	// Get the repository to get its ID
	var repo model.Repository
	if err := db.Where("name = ?", repoName).First(&repo).Error; err != nil {
		return fmt.Errorf("repository not found: %w", err)
	}

	token := os.Getenv("GITHUB_TOKEN")
	org := os.Getenv("GITHUB_ORG")

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls?state=all", org, repoName)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{Timeout: 100 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var rawPRs []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rawPRs); err != nil {
		return err
	}

	for _, pr := range rawPRs {
		number := int(pr["number"].(float64))
		title := pr["title"].(string)
		state := pr["state"].(string)
		createdAt := pr["created_at"].(string)

		var mergedAt *string
		if pr["merged_at"] != nil {
			merged := pr["merged_at"].(string)
			mergedAt = &merged
		}

		userLogin := pr["user"].(map[string]interface{})["login"].(string)

		// Assignees
		var assigneeIDs []*model.Developer
		for _, a := range pr["assignees"].([]interface{}) {
			user := a.(map[string]interface{})
			dev := model.Developer{
				Name:  user["login"].(string),
				Email: fmt.Sprintf("%s@unknown.com", user["login"].(string)), // fallback
			}
			db.FirstOrCreate(&dev, model.Developer{Email: dev.Email})
			assigneeIDs = append(assigneeIDs, &dev)
		}

		// Reviewers
		var reviewerIDs []*model.Developer
		for _, r := range pr["requested_reviewers"].([]interface{}) {
			user := r.(map[string]interface{})
			dev := model.Developer{
				Name:  user["login"].(string),
				Email: fmt.Sprintf("%s@unknown.com", user["login"].(string)), // fallback
			}
			db.FirstOrCreate(&dev, model.Developer{Email: dev.Email})
			reviewerIDs = append(reviewerIDs, &dev)
		}

		var prRecord model.PullRequest
		result := db.Where("pr_number = ? AND repository_id = ?", number, repo.ID).First(&prRecord)
		if result.Error != nil && result.Error.Error() != "record not found" {
			log.Println("DB error fetching PR:", result.Error)
			continue
		}

		prRecord.PRNumber = number
		prRecord.Title = title
		prRecord.State = state
		prRecord.CreatedAt = createdAt
		prRecord.MergedAt = mergedAt
		prRecord.UserLogin = userLogin
		prRecord.RepositoryID = repo.ID
		prRecord.Assignees = assigneeIDs
		prRecord.Reviewers = reviewerIDs

		if result.RowsAffected == 0 {
			err = db.Create(&prRecord).Error
		} else {
			err = db.Save(&prRecord).Error
		}

		if err != nil {
			log.Println("Error saving PR:", err)
		} else {
			log.Printf("PR #%d saved/updated\n", number)
		}
	}

	return nil
}
