package fruit

import "github.com/robertotambunan/gdg-jakarta-2025-demo/entity"

// Repository defines the Elasticsearch operations we need.
type Repository interface {
	SearchFruits(query string) ([]entity.Fruit, error)
}
