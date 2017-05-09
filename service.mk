# Service specific Makefile

include proto.mk

DB_USER ?= root
DB_PASS ?=
DB_HOST ?= 127.0.0.1
DB_PORT ?= 3306
DB_NAME ?= service

.PHONY: migrate

migrate: ## Run migrations
ifdef GODOTENV
	@${GODOTENV} $(MAKE) GODOTENV="" migrate
else
	migrate -path ${PWD}/migrations/ -database mysql://${DB_USER}:${DB_PASS}@tcp\(${DB_HOST}:${DB_PORT}\)/${DB_NAME} up
endif

docker-compose.override.yml: ## Create docker-compose override file
	cp docker-compose.override.yml.example docker-compose.override.yml

setup::
	mkdir -p var/

clean::
	rm -rf var/ docker-compose.override.yml

envcheck::
	$(call executable_check,Docker Compose,docker-compose)
