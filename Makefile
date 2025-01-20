generate-docs:
	cd rest && swag init --parseDependency --parseInternal --parseDepth 1 && mv ./docs ./controllers

run:
	docker compose up --build

test:
	make generate-docs
	docker compose up db -d
	cd rest && \
		POSTGRES_HOST=localhost POSTGRES_PORT=5432 POSTGRES_USER=user \
		POSTGRES_PASSWORD=password POSTGRES_DATABASE=db GIN_MODE=release \
		go test -v ./tests/...
