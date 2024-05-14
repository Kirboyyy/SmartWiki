package vectorstorage

import "github.com/kirboyyy/smartwiki/internal/domain"

type Client interface {
	SimilaritySearch(request domain.SearchRequest) ([]domain.SearchResult, error)
}
