package sqlite

import (
	"database/sql"
	"errors"
	"github.com/kirboyyy/smartwiki/internal/adapters/embedding"
	"github.com/kirboyyy/smartwiki/internal/domain"
)

type Client struct {
	DB              *sql.DB
	EmbeddingClient embedding.Client
}

func (c Client) SimilaritySearch(request domain.SearchRequest) ([]domain.SearchResult, error) {
	_, err := c.EmbeddingClient.Embed(request.Query)
	if err != nil {
		return nil, err
	}
	// TODO: Query DB with value of EmbeddingClient.Embed()
	return nil, errors.New("not implemented yet")
}
