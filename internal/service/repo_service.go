package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Repo struct {
	Name string `json:"name"`
	URL  string `json:"html_url"`
}

func FetchRepositories() ([]Repo, error) {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	token := os.Getenv("GITHUB_TOKEN")
	org := os.Getenv("GITHUB_ORG")

	url := fmt.Sprintf("https://api.github.com/orgs/%s/repos?type=all&per_page=100", org)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	log.Println("Request prepared:", req.URL.String())
	client := &http.Client{
		Timeout: 100 * time.Second, // Timeout after 10 seconds
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Request failed with error:", err)

		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var repos []Repo
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}
