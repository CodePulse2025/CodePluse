// internal/model/task.go
package model

import "time"

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Developer struct {
	ID    uint `gorm:"primaryKey;autoIncrement"` // Local unique ID
	Name  string
	Email string `gorm:"unique"`
}

type Commit struct {
	ID            string `gorm:"primaryKey"` // Commit SHA
	CommitterID   uint   // Foreign key to Developer.ID
	RepoName      string
	CommitMessage string
	CommitHash    string
	CreatedAt     string
	RepositoryID  uint
	Developer     Developer  `gorm:"foreignKey:CommitterID"`
	Repository    Repository `gorm:"foreignKey:RepositoryID"`
}

type PullRequest struct {
	ID           uint `gorm:"primaryKey"`
	PRNumber     int  `gorm:"index"`
	Title        string
	State        string
	CreatedAt    string
	MergedAt     *string
	UserLogin    string
	RepositoryID uint
	Assignees    []*Developer `gorm:"many2many:pr_assignees;"`
	Reviewers    []*Developer `gorm:"many2many:pr_reviewers;"`
}

// Review represents the code review data model
type Review struct {
	ID           string `gorm:"primaryKey"`
	DeveloperID  string
	PRID         string
	ReviewStatus string
	CreatedAt    string
}

type Repository struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	URL       string `gorm:"not null"`
	CreatedAt time.Time
}

type TopContributor struct {
	Name  string
	Email string
	Count int
}
type TopReviewer struct {
	Name        string `json:"name"`
	ReviewCount int    `json:"review_count"`
}
