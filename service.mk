.PHONY: docker-local migrate

docker-compose.override.yml: ## Create docker-compose override file
	cp docker-compose.override.yml.example docker-compose.override.yml

docker-local: docker-compose.override.yml ## Setup local docker env
	mkdir -p var/
	docker-compose up -d

migrate: ## Run migrations
	migrate -path ./migrations/ -database mysql://root:@tcp\(127.0.0.1:3336\)/service up

setup:: docker-compose.override.yml

clean::
	rm -rf docker-compose.override.yml

envcheck::
	$(call executable_check,Docker Compose,docker-compose)
