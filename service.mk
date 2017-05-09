# Service specific Makefile

include proto.mk

DB_USER ?= root
DB_PASS ?=
DB_HOST ?= 127.0.0.1
DB_PORT ?= 3336
DB_NAME ?= service

.PHONY: start migrate

docker-compose.override.yml: ## Create docker-compose override file
	cp docker-compose.override.yml.example docker-compose.override.yml

start: docker-compose.override.yml ## Start docker env
	mkdir -p var/
	docker-compose up -d

migrate: ## Run migrations
	migrate -path ${PWD}/migrations/ -database mysql://${DB_USER}:${DB_PASS}@tcp\(${DB_HOST}:${DB_PORT}\)/${DB_NAME} up

setup:: docker-compose.override.yml

clean::
	docker-compose down
	rm -rf var/ docker-compose.override.yml

envcheck::
	$(call executable_check,Docker Compose,docker-compose)
