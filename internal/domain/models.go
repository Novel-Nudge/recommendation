package domain

type ActionType int

const (
	ActionLikeDB    = 1
	ActionReadDB    = 0
	ActionDislikeDB = -1
)

type NearestBook struct {
	BookID    string
	Distance  float64 // The distance between the user embed and the book embedding
	Embedding []float64
}

type NearestUsers struct {
	UserIDs []string // list of user ID's

	Mapping map[string]float64 // mapping of UserIDs to distance
}

type UserAction struct {
	BookID string
	UserID string

	FinalScore float64

	// Expectation scores for user-book and book-user
	ContentRank float64
	CollabRank  float64

	// What action the user took swiped (1) not (-1)  read (0)
	Action ActionType

	OverlappingTopics []string // Main topics or themes

	BookVector []float64
}

type Features struct {

	// Amount of times we've updated the user features
	Updates int

	// Core traits
	Complexity float32 // 0-1: How complex/sophisticated the content is
	Emotion    float32 // 0-1: Level of emotional intensity
	Pace       float32 // 0-1: How fast-paced vs contemplative

	// Genre weights
	LiteraryFiction   float32
	Mystery           float32
	Romance           float32
	ScienceFiction    float32
	Fantasy           float32
	Thriller          float32
	HistoricalFiction float32
	Contemporary      float32

	// Mood weights
	Uplifting   float32
	Dark        float32
	Mysterious  float32
	Romantic    float32
	Humorous    float32
	Serious     float32
	Melancholic float32
	Inspiring   float32

	// Additional characteristics
	ReadingLevel    float32 // 0-1: Required reading comprehension level
	VocabularyLevel float32 // 0-1: Vocabulary difficulty
	Experimental    float32 // 0-1: How experimental/conventional the writing is
}

type FinalRecommendation struct {
	BookID      string
	FinalScore  float64
	ContentRank float64
	CollabRank  float64
}
