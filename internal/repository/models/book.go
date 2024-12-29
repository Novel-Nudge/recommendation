package models

type BookFeatsDB struct {

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

type NearestBookDB struct {
	ID        string
	Distance  float64 // The distance between the user embed and the book embedding
	Embedding []float64
}

type BookDB struct {
	ID            string
	Title         string
	Author        string
	YearPublished int

	Description string   // Book blurb or description
	Topics      []string // Main topics or themes

}
