// internal/service/metrics/top_reviewers.go
package metrics

import (
	"github.com/hiteshwagh9383/CodePulse/internal/model"
	"github.com/hiteshwagh9383/CodePulse/internal/repository"
)

// GetTopReviewers fetches the top 3 reviewers based on the PR review count
func GetTopReviewers() ([]model.TopReviewer, error) {
	db := repository.GetDB()
	var reviewers []model.TopReviewer

	err := db.Table("pr_reviewers").
		Select("developers.name, COUNT(pr_reviewers.developer_id) AS review_count").
		Joins("JOIN developers ON developers.id = pr_reviewers.developer_id").
		Group("developers.name").
		Order("review_count DESC").
		Limit(3).
		Scan(&reviewers).Error

	if err != nil {
		return nil, err
	}

	return reviewers, nil
}
