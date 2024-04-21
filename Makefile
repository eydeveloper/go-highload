include .env

.DEFAULT_GOAL := help

help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

postgres-master-bash: ## Connect to the PostgreSQL master service using bash.
	docker compose exec -it postgres-master bash

postgres-slave-1-bash: ## Connect to the first PostgreSQL slave service using bash.
	docker compose exec -it postgres-slave-1 bash

postgres-slave-2-bash: ## Connect to the second PostgreSQL slave service using bash.
	docker compose exec -it postgres-slave-2 bash

postgres-master-psql: ## Log in to the PostgreSQL master console from default user
	docker compose exec -it postgres-master psql -U postgres

postgres-slave-1-psql: ## Log in to the first PostgreSQL slave console from default user
	docker compose exec -it postgres-slave-1 psql -U postgres

postgres-slave-2-psql: ## Log in to the second PostgreSQL slave console from default user
	docker compose exec -it postgres-slave-2 psql -U postgres

up: ## Up Docker-project
	docker compose up -d

down: ## Down Docker-project
	docker compose down --remove-orphans

stop: ## Stop Docker-project
	docker compose stop

build: ## Build Docker-project
	docker compose build --no-cache

ps: ## Show list containers
	docker compose ps

default: help
