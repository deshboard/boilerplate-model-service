# Service specific Makefile

include proto.mk

DATABASE_PORT ?= 3336

.PHONY: start migrate

docker-compose.override.yml: ## Create docker-compose override file
	cp docker-compose.override.yml.example docker-compose.override.yml

start: docker-compose.override.yml ## Start docker env
	mkdir -p var/
	docker-compose up -d

# TODO: fix path when migrate is released
migrate: ## Run migrations
	migrate -path ${PWD}/migrations/ -database mysql://root:@tcp\(127.0.0.1:${DATABASE_PORT}\)/service up

setup:: docker-compose.override.yml

clean::
	docker-compose down
	rm -rf var/ docker-compose.override.yml

envcheck::
	$(call executable_check,Docker Compose,docker-compose)
