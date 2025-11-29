package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/robertotambunan/gdg-jakarta-2025-demo/entity"
)

// internal type mirroring the ES search response for our use-case.
type searchResult struct {
	Hits struct {
		Hits []struct {
			Source struct {
				Nama     string  `json:"nama"`
				Kategori string  `json:"kategori"`
				Harga    float64 `json:"harga"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

// SearchFruits performs an autocomplete-style search against ES.
func (r *repository) SearchFruits(query string) ([]entity.Fruit, error) {
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"nama": map[string]interface{}{
					"query":         query,
					"fuzziness":     "AUTO",
					"prefix_length": 2,
				},
			},
		},
		"size": 10,
	}

	jsonData, err := json.Marshal(searchQuery)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/_search", r.baseURL, r.index)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("elasticsearch error: %s", string(body))
	}

	var result searchResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	fruits := make([]entity.Fruit, 0, len(result.Hits.Hits))
	for _, hit := range result.Hits.Hits {
		fruits = append(fruits, entity.Fruit{
			Nama:     hit.Source.Nama,
			Kategori: hit.Source.Kategori,
			Harga:    hit.Source.Harga,
		})
	}

	return fruits, nil
}
