package service

import (
	"github.com/kirboyyy/smartwiki/internal/adapters/vectorstorage"
	"github.com/kirboyyy/smartwiki/internal/domain"
)

type SearchService struct {
	VectorStorageClient vectorstorage.Client
}

func (ss SearchService) searchEmbeddings(request domain.SearchRequest) ([]domain.SearchResult, error) {
	result, err := ss.VectorStorageClient.SimilaritySearch(request)
	if err != nil {
		return nil, err
	}
	return result, nil

}
