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
	docker-compose exec mysql mysql -uroot -puser
start:
	docker-compose exec app go run ./cmd/main.go
down:
	docker-compose down
stop:
	docker-compose stop


.PHONY: build