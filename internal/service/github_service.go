package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Repo struct {
	Name string `json:"name"`
	URL  string `json:"html_url"`
}

func FetchRepositories() ([]Repo, error) {
	token := os.Getenv("GITHUB_TOKEN")
	org := os.Getenv("GITHUB_ORG")

	url := fmt.Sprintf("https://api.github.com/orgs/%s/repos", org)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var repos []Repo
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}
