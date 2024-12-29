package mappers

import "recommendation/internal/repository/models"

// CreateUserMapping Creates an instance of a hash map of userID to float and a list of string for creating a collab match
func CreateUserMapping(users []models.NearestUserDB) (map[string]float64, []string) {
	userList := make([]string, len(users))
	userMap := make(map[string]float64, len(users))

	for i, user := range users {
		userList[i] = user.ID
		userMap[user.ID] = user.Distance
	}

	return userMap, userList
}
