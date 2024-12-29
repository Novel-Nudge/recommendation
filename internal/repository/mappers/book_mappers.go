package mappers

import (
	"recommendation/internal/repository/models"
	"recommendation/pkg/recommendation"
)

func MapContFilterPkg(books []models.NearestBookDB) []recommendation.ContentMatch {
	pkgBooks := make([]recommendation.ContentMatch, len(books))

	for i, book := range books {
		pkgBooks[i] = recommendation.ContentMatch{
			BookID:    book.ID,
			Distance:  book.Distance,
			Embedding: book.Embedding,
		}
	}
	return pkgBooks
}
