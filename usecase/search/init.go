package search

import (
	"context"

	"github.com/robertotambunan/gdg-jakarta-2025-demo/entity"
	"github.com/robertotambunan/gdg-jakarta-2025-demo/repository/fruit"
)

type usecase struct {
	repo fruit.Repository
}

// NewUsecase wires a repository into a search use-case.
func NewUsecase(repo fruit.Repository) Usecase {
	return &usecase{repo: repo}
}

// Autocomplete currently delegates directly to the repository.
// This is where you would add validation, logging, metrics, etc.
func (u *usecase) Autocomplete(ctx context.Context, query string) ([]entity.Fruit, error) {
	return u.repo.SearchFruits(query)
}
