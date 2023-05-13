# Run project
.PHONY: dev
dev:
	@docker compose run --rm api sh -c 'go run cmd/serverd/main.go'

# Setup project
.PHONY: pg pg-migrate-up pg-migrate-down gen-mocks boiler-gen
build-img:
	@docker build -f docker/dev.Dockerfile -t alchemist_api-dev:latest .

pg:
	@docker compose up -d pg

pg-migrate-up:
	@docker compose run --rm pgmigrate sh -c 'migrate -path /migrations -database "$$DB_URL" up'

pg-migrate-down:
	@docker compose run --rm pgmigrate sh -c 'migrate -path /migrations -database "$$DB_URL" drop'

gen-mocks:
	@echo "Mockery generating..."

boiler-gen:
	@docker compose run --rm api sh -c 'sqlboiler psql'
