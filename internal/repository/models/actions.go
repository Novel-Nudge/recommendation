package models

type UserActionDB struct {
	BookID string
	UserID string

	FinalScore float64

	// Expectation scores for user-book and book-user
	ContentRank float64
	CollabRank  float64

	// What action the user took swiped (1) not (-1)  read (0)
	Action int

	OverlappingTopics []string // Main topics or themes

	BookEmbedding []float64
}
