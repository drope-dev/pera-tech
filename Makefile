.PHONY: test run migrate-up postgres-up postgres-down

test:
	go test ./...

run:
	go run ./cmd/api

migrate-up:
	go run ./cmd/migrate

postgres-up:
	docker compose up -d postgres

postgres-down:
	docker compose down
