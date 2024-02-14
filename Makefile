include .env

init:
	@make build
	@make up
build:
	docker-compose build --no-cache
up:
	docker-compose up -d
app:
	docker-compose exec app sh
mysql:
	docker-compose exec mysql mysql -u$(API_DB_USER) -p$(API_DB_PASS)
migrate:
	docker-compose exec app goose -dir ./build/db/migration mysql "$(API_DB_USER):$(API_DB_PASS)@tcp(mysql:$(API_DB_PORT))/$(API_DB_NAME)" up
roll-back:
	docker-compose exec app goose -dir ./build/db/migration mysql "$(API_DB_USER):$(API_DB_PASS)@tcp(mysql:$(API_DB_PORT))/$(API_DB_NAME)" reset
create-migration: # ファイル名は適宜変更すること
	docker-compose exec app goose -dir ./build/db/migration create insert_users sql
create-mock: # ファイル名は適宜変更すること
	docker-compose exec app mockgen -source=pkg/domain/repository/user_repository.go -destination pkg/lib/mock/mock_user.go
start:
	docker-compose exec app go run ./cmd/main.go
down:
	docker-compose down
stop:
	docker-compose stop
gqlgen:
	docker-compose exec app go run github.com/99designs/gqlgen generate
air:
	docker-compose exec app air -c .air.toml
app-dlv:
	docker-compose exec app dlv debug ./cmd/main.go
dlv:
	docker-compose exec app dlv debug ./cmd/main.go --headless --listen=:2345 --api-version=2 --accept-multiclient
dry-fix:
	golangci-lint run ./...
fix:
	golangci-lint run --fix
test:
	go test ./... -coverprofile=coverage.out


.PHONY: build