package elasticsearch

import (
	"net/http"

	"github.com/robertotambunan/gdg-jakarta-2025-demo/repository/fruit"
)

type repository struct {
	baseURL string
	index   string
	client  *http.Client
}

// NewRepository creates a new Elasticsearch-backed repository.
func NewRepository(baseURL, index string) fruit.Repository {
	return &repository{
		baseURL: baseURL,
		index:   index,
		client:  &http.Client{},
	}
}
