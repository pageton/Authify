BINARY_NAME := authfiy
MAIN_FILE := cmd/main.go
MIGRATION_FILE := db/migrations/db_migrations/setup.go
SQLC_CONFIG := db/migrations/sqlc.yaml
DOCKER_IMAGE := authfiy-app
DOCKER_CONTAINER := authfiy-container
DB_CONTAINER := authfiy-db
PORT := $(shell grep ^PORT= .env | cut -d ':' -f2)
NETWORK := host

.PHONY: run migrate build clean rebuild db-up all help docker-build docker-run docker-clean docker-restart docker-compose-up

run: ## Run the project
	go run $(MAIN_FILE)

migrate db-up: ## Run database migrations
	@echo "Running database migrations..."
	go run $(MIGRATION_FILE)

sqlc-generate: ## Generate Go code from SQL queries using sqlc
	sqlc generate -f $(SQLC_CONFIG)

build: ## Build the project (compile)
	go build -o $(BINARY_NAME) $(MAIN_FILE)

clean: ## Clean up the project (remove compiled files)
	rm -f $(BINARY_NAME)

rebuild: clean build ## Rebuild the project

all: build run ## Build and run the project

docker-build: ## Build the Docker image
	docker build -t $(DOCKER_IMAGE) .

docker-run: ## Run the Docker container with logging and network options
	docker run --name $(DOCKER_CONTAINER) --network=$(NETWORK) -p $(PORT):$(PORT) -d $(DOCKER_IMAGE) \
		&& docker logs -f $(DOCKER_CONTAINER)

docker-clean: ## Stop and remove the Docker container
	-docker stop $(DOCKER_CONTAINER)
	-docker rm $(DOCKER_CONTAINER)

docker-logs: ## Fetch logs from the running container
	docker logs $(DOCKER_CONTAINER)

docker-compose-logs: ## Fetch logs from docker-compose services
	docker-compose logs -f

docker-stop: ## Stop the running Docker container without removing it
	docker stop $(DOCKER_CONTAINER)

docker-compose-stop: ## Stop all docker-compose services without removing them
	docker-compose stop

docker-compose-down: ## Shut down and remove all running Docker Compose services, including networks and volumes
	docker-compose down

docker-remove-images: ## Remove all docker images
	docker rmi -f $(docker images -q)

docker-restart: docker-clean docker-run ## Restart the Docker container

docker-compose-up: ## Run with docker-compose
	docker-compose up -d

help: ## Show available targets in Makefile
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
