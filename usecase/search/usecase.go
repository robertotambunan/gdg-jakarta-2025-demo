package search

import (
	"context"

	"github.com/robertotambunan/gdg-jakarta-2025-demo/entity"
)

// Usecase defines the business logic for search/autocomplete.
type Usecase interface {
	Autocomplete(ctx context.Context, query string) ([]entity.Fruit, error)
}
