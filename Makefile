.PHONY: help up down build restart logs clean ps stop start

# Default target
.DEFAULT_GOAL := help

# Variables
COMPOSE_FILE := docker-compose.yml

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

up: ## Start all services in detached mode
	docker-compose -f $(COMPOSE_FILE) up -d
	@echo "Services started. Access:"
	@echo "  - Elasticsearch: http://localhost:9200"
	@echo "  - Cerebro: http://localhost:9000"
	@echo "  - Web App: http://localhost:8081"

down: ## Stop and remove all services
	docker-compose -f $(COMPOSE_FILE) down

stop: ## Stop all services (without removing)
	docker-compose -f $(COMPOSE_FILE) stop

start: ## Start all services (if already created)
	docker-compose -f $(COMPOSE_FILE) start

restart: ## Restart all services
	docker-compose -f $(COMPOSE_FILE) restart

build: ## Build all images
	docker-compose -f $(COMPOSE_FILE) build

rebuild: ## Rebuild all images without cache
	docker-compose -f $(COMPOSE_FILE) build --no-cache

logs: ## Show logs from all services
	docker-compose -f $(COMPOSE_FILE) logs -f

logs-es: ## Show Elasticsearch logs
	docker-compose -f $(COMPOSE_FILE) logs -f elasticsearch

logs-cerebro: ## Show Cerebro logs
	docker-compose -f $(COMPOSE_FILE) logs -f cerebro

logs-web: ## Show Web app logs
	docker-compose -f $(COMPOSE_FILE) logs -f web

ps: ## Show status of all services
	docker-compose -f $(COMPOSE_FILE) ps

clean: ## Stop services and remove volumes (WARNING: deletes data)
	docker-compose -f $(COMPOSE_FILE) down -v
	@echo "All services and volumes removed"

clean-all: clean ## Remove everything including images
	docker-compose -f $(COMPOSE_FILE) down -v --rmi all
	@echo "All services, volumes, and images removed"

shell-es: ## Open shell in Elasticsearch container
	docker-compose -f $(COMPOSE_FILE) exec elasticsearch /bin/bash

shell-web: ## Open shell in Web container
	docker-compose -f $(COMPOSE_FILE) exec web /bin/sh

health: ## Check health of all services
	@echo "Checking Elasticsearch..."
	@curl -s http://localhost:9200/_cluster/health | jq '.' || echo "Elasticsearch not responding"
	@echo "\nChecking Cerebro..."
	@curl -s http://localhost:9000 > /dev/null && echo "Cerebro is running" || echo "Cerebro not responding"
	@echo "Checking Web App..."
	@curl -s http://localhost:8081 > /dev/null && echo "Web App is running" || echo "Web App not responding"

