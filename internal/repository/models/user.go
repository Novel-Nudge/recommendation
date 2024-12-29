package models

type UserFeaturesDB struct {

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

// UserDB represents a user's reading preferences and characteristics
type UserDB struct {
	ID       string
	JoinDate int64

	// Reading characteristics
	Openness      float32 // 0-1: Openness to new/challenging content
	ReadingSpeed  float32 // 0-1: How quickly they typically read
	ReadingLevel  float32 // 0-1: Reading comprehension level
	TimeAvailable float32 // Average reading time available (hours/week)

	// Historical data
	FavoriteBooks  []string // IDs of favorite books
	DislikedBooks  []string // IDs of disliked books
	BooksRead      []string // IDs of books read
	AbandonedBooks []string // IDs of books started but not finished

	// Reading goals
	Goals []string // e.g., "learning", "entertainment", "growth"

}

type NearestUserDB struct {
	ID       string
	Distance float64 // The distance between the user embed and the book embedding
}
