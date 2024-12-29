package config

type RecommendationConfig struct {

	// Size of the embeddings
	VectorSize int

	// Number of historical actions to be used
	UpdateLength int

	// Used for scoring and feature upadting
	ContentWeight float64 // Weight given to content-based similarity
	CollabWeight  float64 // Weight given to collaborative filtering

	// Scoring config data
	DistanceDecay   float64 // How quickly similarity score drops with distance
	MinContentScore float64 // Minimum score threshold for content matches
	MinCollabScore  float64 // Minimum score threshold for collaborative matches

	// Feature update config data
	//Base learning rates
	ContentLR   float64 // For content-based features
	CollabLR    float64 // For collaborative-based features
	EmbeddingLR float64 // For embedding-based updates

	// Decay factors
	TimeDecay       float64 // General time-based decay
	ConfidenceDecay float64 // Reduces update magnitude as confidence grows

	// Minimum thresholds
	MinConfidence float64 // Minimum confidence required for updates

	// Regularization
	L2Reg float64 // L2 regularization factor

}
