CONFIG_FILE := config/config.yaml
DB_HOST := 'localhost'
DB_PORT := $(shell yq e '.database.port' $(CONFIG_FILE))
DB_USER := $(shell yq e '.database.user' $(CONFIG_FILE))
DB_PASSWORD := $(shell yq e '.database.password' $(CONFIG_FILE))
DB_NAME := $(shell yq e '.database.name' $(CONFIG_FILE))

run:
	@echo "Running the application"
	@go run cmd/main.go

goose-create:
	@goose -s -dir database/migrations create "$(word 2, $(MAKECMDGOALS))" sql

goose-up:
	@goose -s -dir database/migrations postgres "host=$(DB_HOST) user=$(DB_USER) dbname=$(DB_NAME) password=$(DB_PASSWORD) sslmode=disable" up

goose-down:
	@goose -s -dir database/migrations postgres "host=$(DB_HOST) user=$(DB_USER) dbname=$(DB_NAME) password=$(DB_PASSWORD) sslmode=disable" down

goose-down-to:
	@goose -s -dir database/migrations postgres "host=$(DB_HOST) user=$(DB_USER) dbname=$(DB_NAME) password=$(DB_PASSWORD) sslmode=disable" down-to $(word 2, $(MAKECMDGOALS))

%:
	@: