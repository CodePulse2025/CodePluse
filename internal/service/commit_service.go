package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hiteshwagh9383/CodePulse/internal/model"
	"github.com/hiteshwagh9383/CodePulse/internal/repository"
	// "github.com/joho/godotenv"
)

func FetchCommitsForRepo(repoName string) ([]model.Commit, error) {
	db := repository.GetDB()

	token := os.Getenv("GITHUB_TOKEN")
	org := os.Getenv("GITHUB_ORG")

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", org, repoName)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{Timeout: 100 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rawCommits []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rawCommits); err != nil {
		return nil, err
	}

	var commits []model.Commit
	for _, c := range rawCommits {
		sha := c["sha"].(string)
		commitData := c["commit"].(map[string]interface{})
		message := commitData["message"].(string)
		date := commitData["author"].(map[string]interface{})["date"].(string)

		committerInfo, ok := commitData["committer"].(map[string]interface{})
		if !ok || committerInfo == nil {
			continue // skip if no committer info (can happen)
		}
		committerName := committerInfo["name"].(string)
		committerEmail := committerInfo["email"].(string)

		developer := model.Developer{
			Name:  committerName,
			Email: committerEmail,
		}

		if err := db.FirstOrCreate(&developer, model.Developer{Email: developer.Email}).Error; err != nil {
			return nil, err
		}
		commit := model.Commit{
			ID:            sha,
			CommitMessage: message,
			CommitHash:    sha,
			CreatedAt:     date,
			RepoName:      repoName,
			CommitterID:   developer.ID,
		}

		commits = append(commits, commit)
	}

	return commits, nil
}
