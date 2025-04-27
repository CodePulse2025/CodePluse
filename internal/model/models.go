// internal/model/task.go
package model

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Developer struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	CreatedAt string
}

// Commit represents the commit data model
type Commit struct {
	ID            string `gorm:"primaryKey"`
	DeveloperID   string
	RepoName      string
	CommitMessage string
	CommitHash    string
	CreatedAt     string
}

// PR represents the pull request data model
type PR struct {
	ID          string `gorm:"primaryKey"`
	DeveloperID string
	RepoName    string
	PRNumber    int
	Status      string
	CreatedAt   string
}

// Review represents the code review data model
type Review struct {
	ID           string `gorm:"primaryKey"`
	DeveloperID  string
	PRID         string
	ReviewStatus string
	CreatedAt    string
}
