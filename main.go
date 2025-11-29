package main

import (
	"log"
	"os"

	httpdelivery "github.com/robertotambunan/gdg-jakarta-2025-demo/delivery/http"
	fruitelasticsearch "github.com/robertotambunan/gdg-jakarta-2025-demo/repository/fruit/elasticsearch"
	"github.com/robertotambunan/gdg-jakarta-2025-demo/usecase/search"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	elasticsearchURL := getEnv("ELASTICSEARCH_URL", "http://localhost:9200")
	indexName := getEnv("ELASTICSEARCH_INDEX", "supermarket_buah")
	port := getEnv("PORT", "8081")

	repo := fruitelasticsearch.NewRepository(elasticsearchURL, indexName)
	uc := search.NewUsecase(repo)

	server, err := httpdelivery.NewServer(":"+port, "web-templates", uc)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	log.Printf("Server starting on http://localhost:%s", port)
	log.Printf("Elasticsearch URL: %s", elasticsearchURL)

	if err := server.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
