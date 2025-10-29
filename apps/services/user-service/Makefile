# Load environment variables from .env file
-include cmd/.env
export

# DEPENDENCIES & TOOLS
tidy: ## Tidy and sync Go module dependencies
	@echo "--> Tidying go modules..."
	@go mod tidy

install-tools: ## Install required Go command-line tools
	@echo "--> Installing tools..."
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest


# CODE GENERATION
generate: ## Generate Go code from SQL queries using sqlc amd GraphQL schema using gqlgen
	@echo "--> Running go generate..."
	@go get github.com/99designs/gqlgen@v0.17.81
	@go run scripts/generate.go
	@echo "--> All code generated."


# DATABASE COMMANDS
db-create: ## Create the local database
	@createdb --host=localhost --username=postgres --port=5432 aetherium_user_service_db || true

db-drop: ## Drop the local database
	@dropdb --host=localhost --username=postgres --port=5432 aetherium_user_service_db || true

db-migrate-up: ## Run all forward database migrations
	@migrate -database "$(DATABASE_URL)" -path app/externals/database/migrations up

db-migrate-down: ## Revert the last database migration
	@migrate -database "$(DATABASE_URL)" -path app/externals/database/migrations down

db-create-migration: ## Create new migration files. Ex: make db-create-migration name=add_indexes
	@migrate create -ext sql -dir app/externals/database/migrations -seq $(name)


# APPLICATION COMMANDS
run: ## Run the main application
	@go run ./cmd/main.go