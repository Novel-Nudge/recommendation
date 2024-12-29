package mappers

import (
	"recommendation/internal/repository/models"
	"recommendation/pkg/recommendation"
)

// MapActionsPkg Maps the actions a
func MapActionsCollabPkg(actionsDB []models.UserActionDB, nearestMap map[string]float64) []recommendation.CollabMatch {
	pkgActions := make([]recommendation.CollabMatch, len(actionsDB))

	// Combine the actions and the nearest user maps distance for a collab match
	for i, action := range actionsDB {
		pkgActions[i] = recommendation.CollabMatch{
			BookID:        action.BookID,
			UserDistance:  nearestMap[action.UserID],
			UserBookScore: action.FinalScore,
		}
	}

	return pkgActions
}

// MapActionsPkg Maps the actions db model to the recommendation pkg model
func MapActionsPkg(actionsDB []models.UserActionDB) []recommendation.UserAction {
	pkgActions := make([]recommendation.UserAction, len(actionsDB))

	for i, action := range actionsDB {
		pkgActions[i] = recommendation.UserAction{
			BookID:      action.BookID,
			FinalScore:  action.FinalScore,
			ContentRank: action.ContentRank,
			CollabRank:  action.CollabRank,

			Action:            int(action.Action),
			OverlappingTopics: action.OverlappingTopics,
			BookVector:        action.BookEmbedding,
		}
	}

	return pkgActions
}
