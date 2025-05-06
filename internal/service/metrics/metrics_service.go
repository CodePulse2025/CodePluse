package metrics

import (
	"github.com/hiteshwagh9383/CodePulse/internal/model"
	"github.com/hiteshwagh9383/CodePulse/internal/repository"
)

func GetTopContributors() ([]model.TopContributor, error) {
	db := repository.GetDB()
	var results []model.TopContributor

	err := db.Raw(`
		SELECT 
			developers.name, 
			developers.email, 
			COUNT(pull_requests.id) AS count
		FROM 
			pull_requests
		JOIN 
			developers ON developers.name = pull_requests.user_login
		GROUP BY 
			developers.name, developers.email
		ORDER BY 
			count DESC
		LIMIT 3;
	`).Scan(&results).Error

	return results, err
}
