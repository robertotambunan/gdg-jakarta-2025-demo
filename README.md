# GDG Jakarta 2025 Demo - Autocomplete with Elasticsearch

## About

This project is a demonstration of implementing autocomplete search functionality using Elasticsearch and Go. It features a web application that allows users to search for fruits with real-time autocomplete suggestions. The application showcases how to integrate Elasticsearch as a search engine backend with a Go-based HTTP server, providing fuzzy matching and fast search capabilities.

The demo includes:
- A Go web application with HTTP server for serving the search interface
- Elasticsearch integration for search and autocomplete functionality
- Docker-based setup for easy deployment and development
- A web UI for searching fruits by name, category, and price
- Cerebro integration for Elasticsearch cluster management and monitoring

## Quick Start with Docker

### Option 1: Using Makefile (Easiest)

```bash
# Show all available commands
make help

# Start all services (Elasticsearch, Cerebro, Web App)
make up

# View logs
make logs

# Check service status
make ps

# Stop all services
make down

# Clean everything (including volumes)
make clean
```

### Option 2: Using Docker Compose (Recommended)

```bash
# Start all services
docker-compose up -d

# Check if Elasticsearch is running
curl http://localhost:9200

# Stop all services
docker-compose down
```

### Option 2: Using Dockerfile directly

```bash
# Build the image
docker build -t elasticsearch-autocomplete .

# Run the container
docker run -d \
  --name elasticsearch \
  -p 9200:9200 \
  -p 9300:9300
  -e discovery.type=single-node \
  -e xpack.security.enabled=false \
  -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" \
  elasticsearch-autocomplete

# Check if Elasticsearch is running
curl http://localhost:9200
```

## Access Points

### Elasticsearch
- **HTTP API**: http://localhost:9200
- **Cluster Health**: http://localhost:9200/_cluster/health
- **Node Info**: http://localhost:9200/_nodes

### Cerebro (Elasticsearch Web UI)
- **Web Interface**: http://localhost:9000
- Connect to Elasticsearch at: `http://elasticsearch:9200` (or `http://localhost:9200` if connecting from host)

### Web Application (Go)
- **Web Interface**: http://localhost:8081
- Search interface for fruits with autocomplete

## Running the Web Application

### Prerequisites
- Go 1.21 or later
- Elasticsearch running (via Docker Compose)

### Steps

1. **Start Elasticsearch**:
   ```bash
   docker-compose up -d
   ```

2. **Set up the index** (use `elasticsearch.http` file or run manually):
   - Create index with mapping
   - Bulk insert data

3. **Run the Go web application**:
   ```bash
   go run main.go
   ```

4. **Access the application**:
   - Open browser: http://localhost:8080
   - Search for fruits using the search bar

### Web Application Features
- Simple search interface with title "Cari Buah"
- Autocomplete search with fuzzy matching
- Displays fruit name, category, and price
- Modern, responsive UI

## Notes

- Security is disabled for demo purposes (`xpack.security.enabled=false`)
- JVM heap size is set to 512MB (adjust if needed)
- Data is persisted in a Docker volume
- Port 9300 is also exposed for node communication (if needed)
- Web application runs on port 8081

## Makefile Commands

The Makefile provides convenient shortcuts for common operations:

- `make help` - Show all available commands
- `make up` - Start all services
- `make down` - Stop and remove all services
- `make build` - Build all Docker images
- `make logs` - View logs from all services
- `make logs-es` - View Elasticsearch logs only
- `make logs-web` - View Web app logs only
- `make ps` - Show status of all services
- `make restart` - Restart all services
- `make clean` - Stop services and remove volumes (WARNING: deletes data)
- `make health` - Check health of all services